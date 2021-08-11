package query

import "fmt"

type SmsNotFoundError struct {
	SmsId string
	Err   error
}

func (e SmsNotFoundError) Error() string {
	return fmt.Sprintf("Sms %s not found", e.SmsId)
}

func (e SmsNotFoundError) Unwrap() error {
	return e.Err
}
