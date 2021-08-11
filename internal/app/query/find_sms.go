package query

import (
	"context"
)

type FindSmsHandler struct {
	readModel GetSingleSmsReadModel
}

func NewFindSmsHandler(readModel GetSingleSmsReadModel) FindSmsHandler {
	if readModel == nil {
		panic("nil SMS read model")
	}

	return FindSmsHandler{readModel: readModel}
}

type GetSingleSmsReadModel interface {
	GetSingleSms(ctx context.Context, uuid string) (SingleSms, error)
}

func (h FindSmsHandler) Handle(ctx context.Context, uuid string) (SingleSms, error) {
	sms, err := h.readModel.GetSingleSms(ctx, uuid)
	if err != nil {
		err = SmsNotFoundError{SmsId: uuid, Err: err}
	}

	return sms, err
}
