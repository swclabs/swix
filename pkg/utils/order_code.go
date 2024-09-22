package utils

import (
	"math/rand"
	"time"
)

func GenerateOrderCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderCode := make([]byte, length)

	for i := range orderCode {
		orderCode[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(orderCode)
}
