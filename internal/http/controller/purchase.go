package controller

import (
	"github.com/labstack/echo/v4"
	"swclabs/swipecore/internal/core/service/purchase"
)

type IPurchase interface {
	AddToCarts(e echo.Context) error
}

type Purchase struct {
	services purchase.IPurchaseService
}

func (purchase Purchase) AddToCarts(e echo.Context) error {
	//TODO implement me
	panic("implement me")
}

var _ IPurchase = (*Purchase)(nil)
