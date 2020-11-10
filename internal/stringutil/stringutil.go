package stringutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterRunes = "abcdefghijklmnopqrstuvwxyz0123456789"

// Random return a random string
func Random(n int, source []rune) string {
	if len(source) == 0 {
		source = []rune(letterRunes)
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = source[rand.Intn(len(source))] // nolint:gosec
	}

	return string(b)
}
