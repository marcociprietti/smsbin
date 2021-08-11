package sms

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	SmsUuid string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("SMS %s not found", e.SmsUuid)
}

type Repository interface {
	Save(ctx context.Context, sms Sms) error
	Find(ctx context.Context, uuid string) (Sms, error)
}
