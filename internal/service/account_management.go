package service

import (
	"errors"
	"fmt"
	"mime/multipart"

	"swclabs/swiftcart/internal/model"
	"swclabs/swiftcart/internal/repo"
	"swclabs/swiftcart/internal/schema"
	"swclabs/swiftcart/pkg/cloud"
	"swclabs/swiftcart/pkg/utils"
	"swclabs/swiftcart/pkg/x/jwt"
)

type AccountManagement struct {
	user    repo.IUsers
	account repo.IAccounts
}

func NewAccountManagement() IAccountManagement {
	return &AccountManagement{
		user:    repo.NewUsers(),
		account: repo.NewAccounts(),
	}
}

func (accountManagement *AccountManagement) SignUp(req *schema.SignUpRequest) error {
	hash, err := jwt.GenPassword(req.Password)
	if err != nil {
		return err
	}
	if err := accountManagement.user.Insert(&model.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "", // TODO:
	}); err != nil {
		return err
	}
	userInfo, err := accountManagement.user.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	return accountManagement.account.Insert(&model.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "swc",
	})
}

func (accountManagement *AccountManagement) Login(req *schema.LoginRequest) (string, error) {
	account, err := accountManagement.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := jwt.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return jwt.GenerateToken(req.Email)
}

func (accountManagement *AccountManagement) UserInfo(email string) (*schema.UserInfo, error) {
	return accountManagement.user.Info(email)
}

func (accountManagement *AccountManagement) UpdateUserInfo(req *schema.UserUpdate) error {
	return accountManagement.user.SaveInfo(&model.User{
		UserID:      req.Id,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})
}

func (accountManagement *AccountManagement) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UpdateImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	return accountManagement.user.SaveInfo(&model.User{
		Email: email,
		Image: resp.SecureURL,
	})
}

func (accountManagement *AccountManagement) OAuth2SaveUser(req *schema.OAuth2SaveUser) error {
	hash, err := jwt.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	if err := accountManagement.user.OAuth2SaveInfo(&model.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       req.Image,
	}); err != nil {
		return err
	}
	userInfo, err := accountManagement.user.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	return accountManagement.account.Insert(&model.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
		Type:     "oauth2-google",
	})
}
