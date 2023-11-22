package service

import (
	"errors"
	"fmt"

	"swclabs/swiftcart/internal/model"
	"swclabs/swiftcart/internal/repo"
	"swclabs/swiftcart/internal/schema"
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

func (accountManagement *AccountManagement) ForgetPassword(email string) error {
	panic("not implement")
}

func (accountManagement *AccountManagement) UpdateUserInfo(req *schema.UserUpdate) error {
	panic("not implement")
}

func (accountManagement *AccountManagement) UploadAvatar() error {
	panic("not implement")
}
