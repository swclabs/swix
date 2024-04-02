// Author: Duc Hung Ho @kieranhoo
// Description: account management service implementation

package service

import (
	"context"
	"errors"
	"mime/multipart"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/internal/helper/resolver"
	"swclabs/swipe-api/internal/helper/tasks"
	"swclabs/swipe-api/pkg/tools"
)

type AccountManagement struct {
	tasks.AccountManagement
	user    domain.IUserRepository
	account domain.IAccountRepository
	address domain.IAddressRepository
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
		address: repo.NewAddresses(),
	}
}

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

func (manager *AccountManagement) UserInfo(ctx context.Context, email string) (*domain.UserInfo, error) {
	// get user information
	return manager.user.Info(ctx, email)
}

func (manager *AccountManagement) UpdateUserInfo(ctx context.Context, req *domain.UserUpdate) error {
	return manager.user.SaveInfo(ctx, &domain.User{
		UserID:      req.Id,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

func (manager *AccountManagement) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	// call worker resolver process this tasks
	resolver.ImagePool.Process(resolver.UserImage{
		Email: email,
		File:  file,
	})
	return nil
}

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

func (manager *AccountManagement) CheckLoginEmail(ctx context.Context, email string) error {
	_, err := manager.account.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("account not found: " + email)
	}
	return nil
}

func (manager *AccountManagement) UpdateAddress(ctx context.Context, data *domain.Addresses) error {
	//TODO:
	return manager.address.New(ctx, data)
}
