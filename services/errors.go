package services

import "errors"

var (
	ErrorInvalidCustom       = errors.New("invalid custom url")
	ErrorCustomAlreadyExists = errors.New("given custom url already exists")
)
