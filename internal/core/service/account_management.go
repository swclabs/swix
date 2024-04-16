// Package service
// Author: Duc Hung Ho @kieranhoo
// Description: account management service implementation
//
// Three layer
//		Controller _____
//		|			   |
//		Service _______|___ Domain
//	 	|			   |
//	 	Repository ___|

package service

import (
	"context"
	"errors"
	"log"
	"mime/multipart"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/internal/workers/tasks"
	"swclabs/swipe-api/pkg/cloud"
	"swclabs/swipe-api/pkg/tools"
)

// AccountManagement implement domain.AccountManagementService
type AccountManagement struct {
	tasks.AccountManagement // embedded tasks to call worker consume
	user                    domain.IUserRepository
	account                 domain.IAccountRepository
	address                 domain.IAddressRepository
}

// NewAccountManagement return new AccountManagement instance
func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
		address: repo.NewAddresses(),
	}
}

// SignUp user to access system, return error if exist
func (manager *AccountManagement) SignUp(ctx context.Context, req *domain.SignUpRequest) error {
	// call repository layer
	return manager.user.TransactionSignUp(ctx, &domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "",
	}, req.Password)

}

// Login to system, return token if error not exist
func (manager *AccountManagement) Login(ctx context.Context, req *domain.LoginRequest) (string, error) {
	// get account form email
	account, err := manager.account.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	// compare input password
	if err := tools.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return tools.GenerateToken(req.Email)
}

// UserInfo return user information from Database
func (manager *AccountManagement) UserInfo(ctx context.Context, email string) (*domain.UserInfo, error) {
	// get user information
	return manager.user.Info(ctx, email)
}

// UpdateUserInfo update user information to database
func (manager *AccountManagement) UpdateUserInfo(ctx context.Context, req *domain.UserUpdate) error {
	// call repository layer
	return manager.user.SaveInfo(ctx, &domain.User{
		Id:          req.Id,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

// UploadAvatar upload image to cloud storage and save img url to database
func (manager *AccountManagement) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	// upload image to image cloud storage
	resp, err := cloud.UploadImages(cloud.Connection(), file)
	if err != nil {
		log.Fatal(err)
	}
	// call repository layer to save user
	return repo.NewUsers().SaveInfo(context.TODO(), &domain.User{
		Email: email,
		Image: resp.SecureURL,
	})
}

// OAuth2SaveUser save user use oauth2 protocol
func (manager *AccountManagement) OAuth2SaveUser(ctx context.Context, req *domain.OAuth2SaveUser) error {
	return manager.user.TransactionSaveOAuth2(
		ctx,
		&domain.User{
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Image:       req.Image,
		})
}

// CheckLoginEmail check email already exist in database
func (manager *AccountManagement) CheckLoginEmail(ctx context.Context, email string) error {
	_, err := manager.account.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("account not found: " + email)
	}
	return nil
}

// UploadAddress update user address to database
func (manager *AccountManagement) UploadAddress(ctx context.Context, data *domain.Addresses) error {
	//TODO:
	return manager.address.Insert(ctx, data)
}
