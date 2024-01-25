package delivery

import (
	"github.com/swclabs/swipe-api/internal/http"
)

type IAdapter interface {
	Run(addr string) error
}

type _Adapter struct {
	server http.IServer
}

func NewAdapter() IAdapter {
	adapter := &_Adapter{
		server: http.New(),
	}
	// initialize cron job
	adapter._StartCommonJob()
	adapter.server.Bootstrap(http.CommonModule)
	return adapter
}

// Run: run all services one server
func (adapter *_Adapter) Run(addr string) error {
	adapter.server.Bootstrap(
		http.AccountManagementModule,
		http.ProductManagementModule,
	)
	return adapter.server.Run(addr)
}

type _AccountManagementAdapter struct {
	server http.IServer
}

func NewAccountManagementAdapter() IAdapter {
	account := &_AccountManagementAdapter{
		server: http.New(),
	}
	// initialize cron job
	account._StartAccountManagementJob()
	account.server.Bootstrap(http.CommonModule)
	return account
}

func (account *_AccountManagementAdapter) Run(addr string) error {
	account.server.Bootstrap(http.AccountManagementModule)
	return account.server.Run(addr)
}

type _ProductManagementAdapter struct {
	server http.IServer
}

func NewProductManagementAdapter() IAdapter {
	product := &_ProductManagementAdapter{
		server: http.New(),
	}
	product.server.Bootstrap(http.ProductManagementModule)
	return product
}

func (product *_ProductManagementAdapter) Run(addr string) error {
	product.server.Bootstrap(http.ProductManagementModule)
	return product.server.Run(addr)
}
