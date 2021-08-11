package sms

import "fmt"

type InvalidSmsError struct {
	Message string
}

func (e InvalidSmsError) Error() string {
	return fmt.Sprintf("Invalid SMS: %s", e.Message)
}
