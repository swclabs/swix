// Author: Duc Hung Ho @kieranhoo
// Description: account management service implementation

package service

import (
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

func (manager *AccountManagement) SignUp(req *domain.SignUpRequest) error {
	// call repository layer
	return manager.user.TransactionSignUp(&domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "",
	}, req.Password)

}

func (manager *AccountManagement) Login(req *domain.LoginRequest) (string, error) {
	// get account form email
	account, err := manager.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	// compare input password
	if err := tools.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return tools.GenerateToken(req.Email)
}

func (manager *AccountManagement) UserInfo(email string) (*domain.UserInfo, error) {
	// get user information
	return manager.user.Info(email)
}

func (manager *AccountManagement) UpdateUserInfo(req *domain.UserUpdate) error {
	return manager.user.SaveInfo(&domain.User{
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

func (manager *AccountManagement) OAuth2SaveUser(req *domain.OAuth2SaveUser) error {
	return manager.user.TransactionSaveOAuth2(&domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       req.Image,
	})
}

func (manager *AccountManagement) CheckLoginEmail(email string) error {
	_, err := manager.account.GetByEmail(email)
	if err != nil {
		return errors.New("account not found: " + email)
	}
	return nil
}

func (manager *AccountManagement) UpdateAddress(data *domain.Addresses) error {
	//TODO:
	return manager.address.New(data)
}
