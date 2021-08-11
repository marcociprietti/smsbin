package infrastructure

import (
	"context"
	"github.com/cipma/smsbin/app/sms"
)

type MemorySmsRepository struct {
	smsMap map[string]sms.Sms
}

func NewMemorySmsRepository() MemorySmsRepository {
	return MemorySmsRepository{
		smsMap: make(map[string]sms.Sms),
	}
}

func (m MemorySmsRepository) Save(ctx context.Context, sms sms.Sms) error {
	m.smsMap[sms.Uuid] = sms

	return nil
}

func (m MemorySmsRepository) Find(ctx context.Context, uuid string) (sms.Sms, error) {
	s, ok := m.smsMap[uuid]
	if !ok {
		return sms.Sms{}, sms.NotFoundError{SmsUuid: uuid}
	}

	return s, nil
}

func (m MemorySmsRepository) FindAll(ctx context.Context) ([]sms.Sms, error) {
	smsList := make([]sms.Sms, 0, len(m.smsMap))

	for _, s := range m.smsMap {
		smsList = append(smsList, s)
	}

	return smsList, nil
}
