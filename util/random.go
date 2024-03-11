package util

import (
	"math/rand"
	"time"
)

func init() {
	// Seed the random number generator
	// with the current time
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandomEmail() string {
	return RandomString(10) + "@" + RandomString(3) + ".com"
}