// Package utils provides utils functionality
package utils

import (
	"sync"

	"swclabs/swix/internal/config"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// BaseSessions is the base session name
const BaseSessions = "swipe_session"

var store *sessions.CookieStore
var lock = &sync.Mutex{}

// NewSession creates a new session
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

// SaveSession saves session
func SaveSession(c echo.Context, sessionName string, key string, value string) error {
	sess, _ := store.Get(c.Request(), sessionName)
	// sess.Options = &sessions.Options{
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	sess.Values[key] = value
	return sess.Save(c.Request(), c.Response())
}

// Session gets session
func Session(c echo.Context, sessionName, key string) interface{} {
	sess, _ := store.Get(c.Request(), sessionName)
	value := sess.Values[key]
	return value
}
