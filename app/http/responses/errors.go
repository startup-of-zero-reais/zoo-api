package responses

import "errors"

var (
	ErrUnhandledPgError         = errors.New("unknown error at our services")
	ErrUserAlreadyRegistered    = errors.New("user already registered")
	ErrUserNotFound             = errors.New("user not found")
	ErrEnclosureNotFound        = errors.New("enclosure not found")
	ErrCannotImportAnimalAge    = errors.New("cannot import animal with invalid age")
	ErrCannotImportAnimalGender = errors.New("cannot import animal with invalid gender")
)
