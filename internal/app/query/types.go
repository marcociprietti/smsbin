package query

import (
	"github.com/cipma/smsbin/app/sms"
	"time"
)

type SingleSms struct {
	Uuid         string              `json:"uuid"`
	From         string              `json:"from"`
	To           string              `json:"to"`
	Text         string              `json:"text"`
	When         time.Time           `json:"when"`
	FromMetadata PhoneNumberMetadata `json:"from_metadata"`
	ToMetadata   PhoneNumberMetadata `json:"to_metadata"`
	TextMetadata TextMetadata        `json:"text_metadata"`
}

func SingleSmsFromModel(model *sms.Sms) SingleSms {
	return SingleSms{
		Uuid:         model.Uuid,
		From:         model.From,
		To:           model.To,
		Text:         model.Text,
		When:         model.When,
		FromMetadata: phoneNumberMetadataFromModel(&model.FromMetadata),
		ToMetadata:   phoneNumberMetadataFromModel(&model.ToMetadata),
		TextMetadata: textMetadataFromModel(&model.TextMetadata),
	}
}

type PhoneNumberMetadata struct {
	CountryCode         int32  `json:"country_code"`
	NationalNumber      uint64 `json:"national_number"`
	RegionCode          string `json:"region_code"`
	IsValid             bool   `json:"is_valid"`
	InternationalFormat string `json:"international_format"`
	IsAlphanumeric      bool   `json:"is_alphanumeric"`
}

func phoneNumberMetadataFromModel(model *sms.PhoneNumberMetadata) PhoneNumberMetadata {
	return PhoneNumberMetadata{
		CountryCode:         model.CountryCode,
		NationalNumber:      model.NationalNumber,
		RegionCode:          model.RegionCode,
		IsValid:             model.IsValid,
		InternationalFormat: model.InternationalFormatted,
		IsAlphanumeric:      model.IsAlphanumeric,
	}
}

type TextMetadata struct {
	Segments  []Segment `json:"segments"`
	TotalSize int       `json:"total_size"`
}

func textMetadataFromModel(model *sms.TextMetadata) TextMetadata {
	segments := make([]Segment, 0, len(model.Segments))
	for _, s := range model.Segments {
		segments = append(segments, segmentFromModel(&s))
	}

	return TextMetadata{
		Segments:  segments,
		TotalSize: model.TotalSize,
	}
}

type Segment struct {
	Data string `json:"data"`
	Size int    `json:"size"`
}

func segmentFromModel(model *sms.Segment) Segment {
	return Segment{
		Data: model.Data,
		Size: model.Size,
	}
}

type AllSmsItem struct {
	Uuid string    `json:"uuid"`
	From string    `json:"from"`
	To   string    `json:"to"`
	When time.Time `json:"when"`
}

func AllSmsItemFromModel(model *sms.Sms) AllSmsItem {
	return AllSmsItem{
		Uuid: model.Uuid,
		From: model.From,
		To:   model.To,
		When: model.When,
	}
}

type AllSms struct {
	Data []AllSmsItem
}
