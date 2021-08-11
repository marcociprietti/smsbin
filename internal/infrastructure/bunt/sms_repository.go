package bunt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cipma/smsbin/app/query"
	"github.com/cipma/smsbin/app/sms"
	"github.com/tidwall/buntdb"
	"log"
)

type SmsRepository struct {
	client *Client
}

func NewBuntSmsRepository(options ClientOptions) SmsRepository {
	client := NewClient(options)

	return SmsRepository{
		client: client,
	}
}

// SmsRepository interface

func (b SmsRepository) Save(ctx context.Context, sms sms.Sms) error {
	db := b.client.Database()

	return db.Update(func(tx *buntdb.Tx) error {
		smsJson, err := json.Marshal(sms)
		if err != nil {
			return err
		}
		_, _, err = tx.Set(sms.Uuid, string(smsJson), nil)
		if err != nil {
			return err
		}

		return nil
	})
}

func (b SmsRepository) Find(ctx context.Context, uuid string) (sms.Sms, error) {
	db := b.client.Database()

	var appSms sms.Sms

	err := db.View(func(tx *buntdb.Tx) error {
		data, err := tx.Get(uuid)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(data), &appSms)
		if err != nil {
			return err
		}

		return nil
	})

	return appSms, err
}

// AllSmsReadModel interface

func (b SmsRepository) AllSms(_ context.Context, size int, lastUuid string) (query.AllSms, error) {
	db := b.client.Database()

	var appSmsList []query.AllSmsItem

	err := db.View(func(tx *buntdb.Tx) error {
		// Retrieve the last item of the previous iteration
		rawData, err := tx.Get(lastUuid)
		var prevSms sms.Sms

		if err == nil {
			_ = json.Unmarshal([]byte(rawData), &prevSms)
		}

		iterator := func(key, value string) bool {
			if len(appSmsList) >= size {
				return false
			}

			var smsModel sms.Sms
			err := json.Unmarshal([]byte(value), &smsModel)
			if err != nil {
				return false
			}

			if prevSms.Uuid != "" && smsModel.Uuid == prevSms.Uuid {
				// skip this record
				return true
			}

			appSms := query.AllSmsItem{
				Uuid: smsModel.Uuid,
				From: smsModel.From,
				To:   smsModel.To,
				When: smsModel.When,
			}

			appSmsList = append(appSmsList, appSms)

			return true
		}

		if prevSms.Uuid != "" {
			log.Println("Using previous item as pivot")

			err = tx.DescendGreaterThan(IndexWhen, fmt.Sprintf(`{"When": "%s"}`, prevSms.When), iterator)
		} else {
			log.Println("Start getting the newest items")

			err = tx.Descend(IndexWhen, iterator)
		}

		return err
	})

	return query.AllSms{
		Data: appSmsList,
	}, err
}

// GetSingleSmsReadModel interface

func (b SmsRepository) GetSingleSms(ctx context.Context, uuid string) (query.SingleSms, error) {
	db := b.client.Database()

	var appSms query.SingleSms

	err := db.View(func(tx *buntdb.Tx) error {
		data, err := tx.Get(uuid)
		if err != nil {
			return err
		}

		var smsModel sms.Sms
		err = json.Unmarshal([]byte(data), &smsModel)
		if err != nil {
			return err
		}

		appSms = query.SingleSms{
			Uuid:         smsModel.Uuid,
			From:         smsModel.From,
			To:           smsModel.To,
			Text:         smsModel.Text,
			When:         smsModel.When,
			FromMetadata: b.unmarshalPhoneNumber(smsModel.FromMetadata),
			ToMetadata:   b.unmarshalPhoneNumber(smsModel.ToMetadata),
			TextMetadata: b.unmarshalTextMetadata(smsModel.TextMetadata),
		}

		return nil
	})

	return appSms, err
}

func (b SmsRepository) unmarshalPhoneNumber(number sms.PhoneNumberMetadata) query.PhoneNumberMetadata {
	return query.PhoneNumberMetadata{
		CountryCode:         number.CountryCode,
		NationalNumber:      number.NationalNumber,
		RegionCode:          number.RegionCode,
		IsValid:             number.IsValid,
		InternationalFormat: number.InternationalFormatted,
		IsAlphanumeric:      number.IsAlphanumeric,
	}
}

func (b SmsRepository) unmarshalTextMetadata(metadata sms.TextMetadata) query.TextMetadata {
	appTextMetadata := query.TextMetadata{
		Segments:  make([]query.Segment, len(metadata.Segments)),
		TotalSize: metadata.TotalSize,
	}

	for i, s := range metadata.Segments {
		appTextMetadata.Segments[i] = query.Segment{
			Data: s.Data,
			Size: s.Size,
		}
	}

	return appTextMetadata
}
