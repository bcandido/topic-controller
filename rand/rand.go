package rand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	generatedString := make([]byte, length)
	for i := range generatedString {
		generatedString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(generatedString)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
