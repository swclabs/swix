package router

import "github.com/labstack/echo/v4"

type IRouter interface {
	Routers(e *echo.Echo)
}

func New(types string) IRouter {
	switch types {
	case TypeCommon:
		return newCommon()
	case TypeDocs:
		return newDocs()
	case TypeProducts:
		return newProducts()
	case TypePosts:
		return newPosts()
	case TypeAccountManagement:
		return newAccountManagement()
	}
	return newCommon()
}
