package sms

import (
	"time"
)

type Sms struct {
	Uuid         string
	Text         string
	From         string
	To           string
	When         time.Time
	FromMetadata PhoneNumberMetadata
	ToMetadata   PhoneNumberMetadata
	TextMetadata TextMetadata
}

type PhoneNumberMetadata struct {
	CountryCode            int32
	NationalNumber         uint64
	RegionCode             string
	IsValid                bool
	InternationalFormatted string
	IsAlphanumeric         bool
}

type TextMetadata struct {
	Segments  []Segment
	TotalSize int
}

type Segment struct {
	Data string
	Size int
}

func NewSms(uuid string, text string, from string, to string, phoneParser PhoneNumberParser, textParser TextSegmentParser) (Sms, error) {
	if uuid == "" {
		return Sms{}, InvalidSmsError{Message: "empty SMS UUID"}
	}
	if from == "" {
		return Sms{}, InvalidSmsError{Message: "empty SMS from"}
	}
	if to == "" {
		return Sms{}, InvalidSmsError{Message: "empty SMS to"}
	}

	return Sms{
		Uuid:         uuid,
		Text:         text,
		From:         from,
		To:           to,
		When:         time.Now(),
		FromMetadata: phoneParser.Parse(from, true),
		ToMetadata:   phoneParser.Parse(to, false),
		TextMetadata: textParser.ExtractSegments(text),
	}, nil
}
