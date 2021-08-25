package storm

import (
	"context"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/cipma/smsbin/app/query"
	"github.com/cipma/smsbin/app/sms"
)

type SmsRepository struct {
	db *storm.DB
}

type Options struct {
	Path string
	Name string
}

func NewSmsRepository(options Options) SmsRepository {
	db, err := storm.Open(fmt.Sprintf("%s/%s", options.Path, options.Name))
	if err != nil {
		panic(err)
	}

	err = db.Init(&sms.Sms{})
	if err != nil {
		panic(err)
	}

	return SmsRepository{
		db: db,
	}
}

func (s SmsRepository) GetSingleSms(ctx context.Context, uuid string) (query.SingleSms, error) {
	var smsModel sms.Sms

	err := s.db.One("Uuid", uuid, &smsModel)
	if err != nil {
		return query.SingleSms{}, err
	}

	appSms := query.SingleSmsFromModel(&smsModel)

	return appSms, err
}

func (s SmsRepository) AllSms(ctx context.Context, size int, lastUuid string) (query.AllSms, error) {
	matchers := make([]q.Matcher, 0)

	prevSms, err := s.Find(ctx, lastUuid)
	if err == nil {
		matchers = append(matchers, q.Lt("When", prevSms.When))
	}

	dbQuery := s.db.Select(matchers...)

	if size > 0 {
		dbQuery = dbQuery.Limit(size)
	}

	var smsList []sms.Sms
	err = dbQuery.OrderBy("When").Reverse().Find(&smsList)
	if err != nil {
		return query.AllSms{}, err
	}

	allSms := query.AllSms{
		Data: make([]query.AllSmsItem, 0, len(smsList)),
	}

	for _, model := range smsList {
		allSms.Data = append(allSms.Data, query.AllSmsItemFromModel(&model))
	}

	return allSms, nil
}

func (s SmsRepository) Save(ctx context.Context, sms sms.Sms) error {
	return s.db.Save(&sms)
}

func (s SmsRepository) Find(ctx context.Context, uuid string) (sms.Sms, error) {
	var smsModel sms.Sms
	err := s.db.One("Uuid", uuid, &smsModel)

	return smsModel, err
}
