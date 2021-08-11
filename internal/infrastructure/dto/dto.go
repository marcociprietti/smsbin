package dto

import (
	"errors"
)

type CreateSms struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type CreateSmsResponse struct {
	Id string `json:"id"`
}

type Error struct {
	Message  string `json:"error"`
	Previous *Error `json:"previous,omitempty"`
}

func NewError(err error) Error {
	errorDto := Error{
		Message: err.Error(),
	}

	if prevErr := errors.Unwrap(err); prevErr != nil {
		prevErrDto := NewError(prevErr)
		errorDto.Previous = &prevErrDto
	}

	return errorDto
}
