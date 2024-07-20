package utils

import (
	"crypto/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n uint) string {
	b := make([]byte, n)

	_, err := rand.Read(b[:])
	if err != nil {
		panic("could not generate a random string")
	}

	for i := range b {
		b[i] = letterBytes[int64(b[i])%int64(len(letterBytes))]
	}

	return string(b)
}
