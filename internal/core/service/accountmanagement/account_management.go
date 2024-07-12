// Package accountmanagement account management service implementation
// Three layer
//
//		Controller_____
//		|			   |
//		Service _______|___ Domain
//	 	|			   |
//	 	Repository ____|
package accountmanagement

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/users"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/jwt"
	"swclabs/swipecore/pkg/utils"

	"swclabs/swipecore/internal/core/domain"
)

// AccountManagement implement domain.AccountManagementService
type AccountManagement struct {
	Blob    blob.IBlobStorage
	User    users.IUserRepository
	Account accounts.IAccountRepository
	Address addresses.IAddressRepository
}

var _ IAccountManagement = (*AccountManagement)(nil)

// New create new AccountManagement object
func New(
	blob blob.IBlobStorage,
	user users.IUserRepository,
	account accounts.IAccountRepository,
	address addresses.IAddressRepository,
) IAccountManagement {
	return &AccountManagement{
		Blob:    blob,
		User:    user,
		Account: account,
		Address: address,
	}
}

// SignUp user to access system, return error if exist
func (manager *AccountManagement) SignUp(ctx context.Context, req domain.SignUpSchema) error {
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)
	if err := userRepo.Insert(ctx,
		domain.Users{
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
	hashPassword, err := jwt.GenPassword(req.Password)
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

	if err := accountRepo.Insert(ctx, domain.Account{
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
func (manager *AccountManagement) Login(
	ctx context.Context, req domain.LoginSchema) (string, error) {
	// get account form email
	account, err := manager.Account.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	// compare input password
	if err := jwt.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return jwt.GenerateToken(req.Email)
}

// UserInfo return user information from Database
func (manager *AccountManagement) UserInfo(
	ctx context.Context, email string) (*domain.UserSchema, error) {
	// get user information
	return manager.User.Info(ctx, email)
}

// UpdateUserInfo update user information to database
func (manager *AccountManagement) UpdateUserInfo(
	ctx context.Context, req domain.UserUpdate) error {
	// call repository layer
	return manager.User.SaveInfo(ctx, domain.Users{
		ID:          req.ID,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

// UploadAvatar upload image to blob storage and save img url to database
func (manager *AccountManagement) UploadAvatar(
	email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	// upload image to image blob storage
	resp, err := manager.Blob.UploadImages(file)
	if err != nil {
		log.Fatal(err)
	}
	// call repository layer to save user
	return manager.User.SaveInfo(context.TODO(), domain.Users{
		Email: email,
		Image: resp.SecureURL,
	})
}

// OAuth2SaveUser save user use oauth2 protocol
func (manager *AccountManagement) OAuth2SaveUser(
	ctx context.Context, req domain.OAuth2SaveUser) error {
	hash, err := jwt.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)
	if err := userRepo.OAuth2SaveInfo(ctx, domain.Users{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       req.Image,
	}); err != nil {
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
	if err := accountRepo.Insert(ctx, domain.Account{
		Username: fmt.Sprintf("user#%d", userInfo.ID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "oauth2-google",
	}); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	return tx.Commit(ctx)
}

// CheckLoginEmail check email already exist in database
func (manager *AccountManagement) CheckLoginEmail(
	ctx context.Context, email string) error {
	_, err := manager.Account.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("account not found: " + email)
	}
	return nil
}

// UploadAddress update user address to database
func (manager *AccountManagement) UploadAddress(
	ctx context.Context, data domain.Addresses) error {
	return manager.Address.Insert(ctx, data)
}
