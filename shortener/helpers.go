package shortener

import (
	"math/rand"
	"net/url"
	"time"
)

const _charset string = "ABCDEFGHIJKLMNOPRSTUVYZWabcdefghijklmnoprstuvyzw0123456789"

var source = rand.NewSource(time.Now().UnixNano())

func generateRand(length int) string {
	x := make([]byte, length)
	for i := range x {
		x[i] = _charset[source.Int63()%int64(len(_charset))]
	}
	return string(x)
}

func isValidURL(testUrl string) bool {
	_, err := url.ParseRequestURI(testUrl)
	return err == nil
}
