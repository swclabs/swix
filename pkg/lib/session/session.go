// Package session provides session functionality
package session

import (
	"swclabs/swix/internal/config"
	"sync"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// Base is the base session name
const Base = "swipe_session"

var store *sessions.CookieStore
var lock = &sync.Mutex{}

// New creates a new session
func New() *sessions.CookieStore {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			store = sessions.NewCookieStore([]byte(config.Secret))
		}
	}
	return store
}

// Save saves session
func Save(c echo.Context, sessionName string, key string, value string) error {
	sess, _ := store.Get(c.Request(), sessionName)
	// sess.Options = &sessions.Options{
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	sess.Values[key] = value
	return sess.Save(c.Request(), c.Response())
}

// Get gets session
func Get(c echo.Context, sessionName, key string) string {
	sess, _ := store.Get(c.Request(), sessionName)
	value := sess.Values[key]
	return value.(string)
}
