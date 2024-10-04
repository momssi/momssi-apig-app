package form

import (
	"errors"
	"fmt"
)

const (
	NoError                int = 0
	ErrParsing                 = 4001
	ErrDuplicatedUsername      = 4002
	ErrMissingToken            = 4101
	ErrInvalidToken            = 4102
	ErrInternalServerError     = 5004
	ErrFailGenerateJWTKey      = 5005
)

var codeToMessage = map[int]error{
	NoError:                nil,
	ErrParsing:             errors.New("invalid request body"),
	ErrDuplicatedUsername:  errors.New("duplicated username"),
	ErrMissingToken:        errors.New("missing token"),
	ErrInvalidToken:        errors.New("invalid token"),
	ErrInternalServerError: errors.New("internal server error"),
	ErrFailGenerateJWTKey:  errors.New("could not generate token"),
}

func GetCustomErrMessage(code int, error string) string {
	message, exists := codeToMessage[code]
	if !exists {
		return "Unknown error"
	}

	return fmt.Sprintf("%s, err : %v", message, error)
}

func GetCustomErr(code int) error {
	err, exists := codeToMessage[code]
	if !exists {
		return errors.New("unknown error")
	}

	return err
}

func GetCustomMessage(code int) string {
	message, exists := codeToMessage[code]
	if !exists {
		return "Unknown error"
	}

	if message == nil {
		return "ok"
	}

	return message.Error()
}
