package query

import (
	"context"
)

type AllSmsQuery struct {
}

type AllSmsHandler struct {
	readModel AllSmsReadModel
}

func NewAllSmsHandler(readModel AllSmsReadModel) AllSmsHandler {
	if readModel == nil {
		panic("nil read model")
	}

	return AllSmsHandler{readModel: readModel}
}

type AllSmsReadModel interface {
	AllSms(ctx context.Context, size int, lastUuid string) (AllSms, error)
}

func (h AllSmsHandler) Handle(ctx context.Context, size int, lastUuid string) (AllSms, error) {
	return h.readModel.AllSms(ctx, size, lastUuid)
}
