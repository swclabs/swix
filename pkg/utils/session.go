package utils

import (
	"sync"

	"swclabs/swipecore/internal/config"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const BaseSessions = "session"

var store *sessions.CookieStore
var lock *sync.Mutex = &sync.Mutex{}

func NewSession() *sessions.CookieStore {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			store = sessions.NewCookieStore([]byte(config.Secret))
		}
	}
	return store
}

func SaveSession(c echo.Context, sessionName string, key string, value string) error {
	sess, _ := store.Get(c.Request(), sessionName)
	// sess.Options = &sessions.Options{
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	sess.Values[key] = value
	return sess.Save(c.Request(), c.Response())
}

func Session(c echo.Context, sessionName, key string) interface{} {
	sess, _ := store.Get(c.Request(), sessionName)
	value := sess.Values[key]
	return value
}
