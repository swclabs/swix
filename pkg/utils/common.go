// Package utils provides utils functionality
package utils

import (
	"math/rand"
	"net/mail"
	"strings"
	"time"
)

// IsEmail checks if the given string is a valid email
func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// RandomString generates a random string of the given length
func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

// HanoiTimezone returns the time in Hanoi timezone
func HanoiTimezone(t time.Time) string {
	return HanoiZone(t).Format(time.DateTime)
}

// HanoiZone returns the time in Hanoi timezone
func HanoiZone(t time.Time) time.Time {
	return t.In(time.FixedZone("GMT+7", 7*60*60))
}

// RemoveSpace removes all spaces in the given string
func RemoveSpace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

