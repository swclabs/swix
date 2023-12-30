// Author: Duc Hung Ho @ikierans
// Description: account management service implementation

package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"swclabs/swiftcart/internal/model"
	"swclabs/swiftcart/internal/repo"
	"swclabs/swiftcart/internal/schema"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/internal/tasks/plugin"
	"swclabs/swiftcart/pkg/jwt"
	"swclabs/swiftcart/pkg/utils"
)

type AccountManagement struct {
	user    repo.IUsers
	account repo.IAccounts
	tasks   *tasks.AccountManagement
}

func NewAccountManagement() IAccountManagement {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
		tasks:   tasks.NewAccountManagement(),
	}
}

func (manager *AccountManagement) SignUp(req *schema.SignUpRequest) error {
	// Add task
	// Begin:
	hash, err := jwt.GenPassword(req.Password)
	if err != nil {
		return err
	}
	if err := manager.user.Insert(&model.User{
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
	return manager.account.Insert(&model.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "swc",
	})
	// End
}

func (manager *AccountManagement) Login(req *schema.LoginRequest) (string, error) {
	account, err := manager.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := jwt.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return jwt.GenerateToken(req.Email)
}

func (manager *AccountManagement) UserInfo(email string) (*schema.UserInfo, error) {
	return manager.user.Info(email)
}

func (manager *AccountManagement) UpdateUserInfo(req *schema.UserUpdate) error {
	// Add task
	// Begin
	return manager.user.SaveInfo(&model.User{
		UserID:      req.Id,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
	// End
}

func (manager *AccountManagement) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	plugin.ImagePool.Process(plugin.Image{
		Email: email,
		File:  file,
	})
	return nil
}

func (manager *AccountManagement) OAuth2SaveUser(req *schema.OAuth2SaveUser) error {
	hash, err := jwt.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	// Add task
	// Begin:
	if err := manager.user.OAuth2SaveInfo(&model.User{
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
	return manager.account.Insert(&model.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "oauth2-google",
	})
	// End
}
