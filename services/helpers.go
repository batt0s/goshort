package services

import (
	"net/url"
	"unicode"
)

func isValidURL(testUrl string) bool {
	_, err := url.ParseRequestURI(testUrl)
	return err == nil
}

func onlyLetterOrDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
