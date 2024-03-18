// Author: Duc Hung Ho @kieranhoo
// Description: account management service implementation

package service

import (
	"errors"
	"mime/multipart"

	"github.com/swclabs/swipe-server/internal/domain"
	"github.com/swclabs/swipe-server/internal/helper/resolver"
	"github.com/swclabs/swipe-server/internal/repo"
	"github.com/swclabs/swipe-server/internal/tasks"
	"github.com/swclabs/swipe-server/pkg/tools"
)

type AccountManagement struct {
	tasks.AccountManagement
	user    domain.IUserRepository
	account domain.IAccountRepository
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
	}
}

func (manager *AccountManagement) SignUp(req *domain.SignUpRequest) error {
	// TODO: update transaction
	return manager.user.SignUp(&domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "",
	}, req.Password)

}

func (manager *AccountManagement) Login(req *domain.LoginRequest) (string, error) {
	account, err := manager.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := tools.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return tools.GenerateToken(req.Email)
}

func (manager *AccountManagement) UserInfo(email string) (*domain.UserInfo, error) {
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
	resolver.ImagePool.Process(resolver.Image{
		Email: email,
		File:  file,
	})
	return nil
}

func (manager *AccountManagement) OAuth2SaveUser(req *domain.OAuth2SaveUser) error {
	return manager.user.SaveOAuth2(&domain.User{
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

func (manager *AccountManagement) UpdateUserAddress(data *domain.Addresses) error {
	panic("implement me")
}
