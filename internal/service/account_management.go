// Author: Duc Hung Ho @kieranhoo
// Description: account management service implementation

package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/internal/repo"
	"swclabs/swiftcart/internal/resolver"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/pkg/jwt"
	"swclabs/swiftcart/pkg/utils"
)

type AccountManagement struct {
	user    domain.IUserRepository
	account domain.IAccountRepository
	Task    tasks.AccountManagementTask
}

func NewAccountManagement() domain.IAccountManagementService {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
	}
}

func (manager *AccountManagement) SignUp(req *domain.SignUpRequest) error {
	hash, err := jwt.GenPassword(req.Password)
	if err != nil {
		return err
	}
	if err := manager.user.Insert(&domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "", // TODO:
	}); err != nil {
		return err
	}
	userInfo, err := manager.user.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	return manager.account.Insert(&domain.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "swc",
	})
}

func (manager *AccountManagement) Login(req *domain.LoginRequest) (string, error) {
	account, err := manager.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := jwt.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return jwt.GenerateToken(req.Email)
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
	hash, err := jwt.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	if err := manager.user.OAuth2SaveInfo(&domain.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       req.Image,
	}); err != nil {
		return err
	}
	userInfo, err := manager.user.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	return manager.account.Insert(&domain.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "oauth2-google",
	})
}
