// Package manager manager service implementation
// Three layer
//
//		Controller_____
//		|			   |
//		Service _______|___ Domain
//	 	|			   |
//	 	Repository ____|
package authentication

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
	"swclabs/swipex/internal/core/repos/accounts"
	"swclabs/swipex/internal/core/repos/addresses"
	"swclabs/swipex/internal/core/repos/users"
	"swclabs/swipex/pkg/infra/blob"
	"swclabs/swipex/pkg/infra/db"
	"swclabs/swipex/pkg/lib/crypto"
	"swclabs/swipex/pkg/utils"

	"github.com/jackc/pgx/v5"
)

var _ IAuthentication = (*Authentication)(nil)
var _ = app.Service(New)

// New create new Authentication object
func New(
	blob blob.IBlobStorage,
	user users.IUsers,
	account accounts.IAccounts,
	address addresses.IAddress,
) IAuthentication {
	return &Authentication{
		Blob:    blob,
		User:    user,
		Account: account,
		Address: address,
	}
}

// Authentication implement IAuthentication
type Authentication struct {
	Blob    blob.IBlobStorage
	User    users.IUsers
	Account accounts.IAccounts
	Address addresses.IAddress
}

// SignUp user to access system, return error if exist
func (auth *Authentication) SignUp(ctx context.Context, req dtos.SignUpRequest) error {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return err
	}

	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)

	if _, err := userRepo.Insert(ctx,
		entity.User{
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Image:       "",
		}); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}

	hashPassword, err := crypto.GenPassword(req.Password)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}

	userInfo, err := userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}

	if _, err = accountRepo.Insert(ctx, entity.Account{
		Username: fmt.Sprintf("user#%d", userInfo.ID),
		Password: hashPassword,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "swc",
	}); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	return tx.Commit(ctx)
}

// Login to system, return token if error not exist
func (auth *Authentication) Login(ctx context.Context, req dtos.LoginRequest) (string, error) {
	// get account form email
	account, err := auth.Account.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	user, err := auth.UserInfo(ctx, req.Email)
	if err != nil {
		return "", err
	}

	// compare input password
	if err := crypto.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return crypto.GenerateToken(user.ID, account.Email, account.Role)
}

// UserInfo return user information from Database
func (auth *Authentication) UserInfo(ctx context.Context, email string) (*model.Users, error) {
	// get user information
	return auth.User.Info(ctx, email)
}

// UpdateUserInfo update user information to database
func (auth *Authentication) UpdateUserInfo(ctx context.Context, req dtos.UserUpdate) error {
	// call repos layer
	return auth.User.Save(ctx, entity.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

// UploadAvatar upload image to blob storage and save img url to database
func (auth *Authentication) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	// upload image to image blob storage
	resp, err := auth.Blob.UploadImages(file)
	if err != nil {
		log.Fatal(err)
	}

	// call repos layer to save user
	return auth.User.Save(context.TODO(), entity.User{
		Email: email,
		Image: resp.SecureURL,
	})
}

// OAuth2SaveUser save user use oauth2 protocol
func (auth *Authentication) OAuth2SaveUser(ctx context.Context, req dtos.OAuth2SaveUser) (userID int64, err error) {
	hash, err := crypto.GenPassword(utils.RandomString(18))
	if err != nil {
		return -1, err
	}

	tx, err := db.NewTx(ctx)
	if err != nil {
		return -1, err
	}

	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)

	if err := userRepo.OAuth2SaveInfo(ctx, entity.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       req.Image,
	}); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return -1, fmt.Errorf("error saving user info: %v", err)
	}

	userInfo, err := userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return -1, err
	}
	
	_, err = accountRepo.Insert(ctx, entity.Account{
		Username: fmt.Sprintf("user#%d", userInfo.ID),
		Password: hash,
		Role:     "customer",
		Email:    req.Email,
		Type:     "oauth2-google",
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return -1, fmt.Errorf("error saving account info: %v", err)
	}
	return userInfo.ID, tx.Commit(ctx)
}

// CheckLoginEmail check email already exist in database
func (auth *Authentication) CheckLoginEmail(ctx context.Context, email string) error {
	account, err := auth.Account.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	if account.Email == email {
		return errors.New("email already exist: " + email)
	}
	return nil
}
