package service

import (
	"errors"
	"example/swiftcart/internal/model"
	"example/swiftcart/internal/repo"
	"example/swiftcart/internal/schema"
	"example/swiftcart/pkg/utils"
	"fmt"
)

type Users struct {
	userInf repo.IUsers
	account repo.IAccounts
}

func NewUsers() IUser {
	return &Users{
		userInf: repo.NewUsers(),
		account: repo.NewAccounts(),
	}
}

func (usr *Users) SignUp(req *schema.SignUpRequest) error {
	hash, err := utils.GenPassword(req.Password)
	if err != nil {
		return err
	}
	if err := usr.userInf.Insert(&model.User{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Image:       "", // TODO:
	}); err != nil {
		return err
	}
	userInfo, err := usr.userInf.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	return usr.account.Insert(&model.Account{
		Username: fmt.Sprintf("user#%d", userInfo.UserID),
		Password: hash,
		Role:     "Customer",
		Email:    req.Email,
	})
}

func (usr *Users) Login(req *schema.LoginRequest) (string, error) {
	account, err := usr.account.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := utils.ComparePassword(account.Password, req.Password); err != nil {
		return "", errors.New("email or password incorrect")
	}
	return utils.GenerateToken(req.Email)
}

func (usr *Users) Infor(email string) (*schema.UserInfor, error) {
	return usr.userInf.Infor(email)
}
