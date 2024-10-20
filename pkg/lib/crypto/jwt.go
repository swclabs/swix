// Package crypto implements the json web token
package crypto

import (
	"errors"
	"fmt"
	"strings"
	"swclabs/swix/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func claims(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.JwtSecret), nil
	})
}

// GenerateToken Generate JWT token
func GenerateToken(userID int64, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	tokenString, errToken := token.SignedString([]byte(config.JwtSecret))
	if errToken != nil {
		return "", errors.New("failed to create token")
	}
	return tokenString, nil
}

// ParseToken Return Email from JWT token
func ParseToken(tokenString string) (userID int64, email string, err error) {
	tokenString = RemoveBearerPrefix(tokenString)
	token, err := claims(tokenString)
	if err != nil {
		return -1, "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email = claims["email"].(string)
		userID = int64(claims["user_id"].(float64))
		// _ = claims["exp"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return -1, "", errors.New("token has expired")
		}
		return userID, email, nil
	}
	return -1, "", errors.New("token invalid")
}

// ParseTokenRole Return Role from JWT token
func ParseTokenRole(tokenString string) (string, error) {
	tokenString = RemoveBearerPrefix(tokenString)
	token, _ := claims(tokenString)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role := claims["role"]
		// _ = claims["exp"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("token has expired")
		}
		if role == nil {
			return "", errors.New("role not found")
		}
		return role.(string), nil
	}
	return "", errors.New("token invalid")
}

func Authenticate(c echo.Context) (userID int64, email string, err error) {
	authHeader := c.Request().Header.Get("Authorization")
	// fmt.Println(authHeader)
	if authHeader == "" || authHeader == "Bearer" {
		return -1, "", errors.New("unauthorized")
	}
	userID, email, err = ParseToken(authHeader)
	if err != nil {
		return -1, "", errors.New("unauthorized")
	}
	return userID, email, nil
}
