package utils

import (
	"math/rand"
	"net/mail"
	"time"
)

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func HanoiTimezone(t time.Time) string {
	return HanoiZone(t).Format(time.DateTime)
}

func HanoiZone(t time.Time) time.Time {
	return t.In(time.FixedZone("GMT+7", 7*60*60))
}
