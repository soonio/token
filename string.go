package token

import (
	"math/rand"
	"time"
)

var seeder = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func String(length int64) string {
	b := make([]byte, length)
	for i := range b {
		randomChar := seeder.Intn(len(charset))
		b[i] = charset[randomChar]
	}
	return string(b)
}
