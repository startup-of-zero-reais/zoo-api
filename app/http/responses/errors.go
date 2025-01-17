package responses

import "errors"

var (
	ErrUnhandledPgError      = errors.New("unknown error at our services")
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrUserNotFound          = errors.New("user not found")
)
