// Package manager manager service implementation
// Three layer
//
//		Controller_____
//		|			   |
//		Service _______|___ Domain
//	 	|			   |
//	 	Repository ____|
package manager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/internal/core/repos/accounts"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/users"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/lib/crypto"
	"swclabs/swix/pkg/utils"
)

var _ IManager = (*Manager)(nil)
var _ = boot.Service(New)

// New create new Manager object
func New(
	blob blob.IBlobStorage,
	user users.IUsers,
	account accounts.IAccounts,
	address addresses.IAddress,
) IManager {
	return &Manager{
		Blob:    blob,
		User:    user,
		Account: account,
		Address: address,
	}
}

// Manager implement IManager
type Manager struct {
	Blob    blob.IBlobStorage
	User    users.IUsers
	Account accounts.IAccounts
	Address addresses.IAddress
}

// SignUp user to access system, return error if exist
func (manager *Manager) SignUp(ctx context.Context, req dtos.SignUpRequest) error {
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)
	if err := userRepo.Insert(ctx,
		entity.Users{
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

	if err := accountRepo.Insert(ctx, entity.Account{
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
func (manager *Manager) Login(ctx context.Context, req dtos.LoginRequest) (string, error) {
	// get account form email
	account, err := manager.Account.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	// compare input password
	if err := crypto.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return crypto.GenerateToken(req.Email, account.Role)
}

// UserInfo return user information from Database
func (manager *Manager) UserInfo(ctx context.Context, email string) (*model.Users, error) {
	// get user information
	return manager.User.Info(ctx, email)
}

// UpdateUserInfo update user information to database
func (manager *Manager) UpdateUserInfo(ctx context.Context, req dtos.UserUpdate) error {
	// call repos layer
	return manager.User.Save(ctx, entity.Users{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

// UploadAvatar upload image to blob storage and save img url to database
func (manager *Manager) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	// upload image to image blob storage
	resp, err := manager.Blob.UploadImages(file)
	if err != nil {
		log.Fatal(err)
	}
	// call repos layer to save user
	return manager.User.Save(context.TODO(), entity.Users{
		Email: email,
		Image: resp.SecureURL,
	})
}

// OAuth2SaveUser save user use oauth2 protocol
func (manager *Manager) OAuth2SaveUser(ctx context.Context, req dtos.OAuth2SaveUser) error {
	hash, err := crypto.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		userRepo    = users.New(tx)
		accountRepo = accounts.New(tx)
	)
	if err := userRepo.OAuth2SaveInfo(ctx, entity.Users{
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
	if err := accountRepo.Insert(ctx, entity.Account{
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
func (manager *Manager) CheckLoginEmail(ctx context.Context, email string) error {
	_, err := manager.Account.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("account not found: " + email)
	}
	return nil
}
