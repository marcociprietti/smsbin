package query

import "time"

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

type PhoneNumberMetadata struct {
	CountryCode         int32  `json:"country_code"`
	NationalNumber      uint64 `json:"national_number"`
	RegionCode          string `json:"region_code"`
	IsValid             bool   `json:"is_valid"`
	InternationalFormat string `json:"international_format"`
	IsAlphanumeric      bool   `json:"is_alphanumeric"`
}

type TextMetadata struct {
	Segments  []Segment `json:"segments"`
	TotalSize int       `json:"total_size"`
}

type Segment struct {
	Data string `json:"data"`
	Size int    `json:"size"`
}

type AllSmsItem struct {
	Uuid string    `json:"uuid"`
	From string    `json:"from"`
	To   string    `json:"to"`
	When time.Time `json:"when"`
}

type AllSms struct {
	Data []AllSmsItem
	Cursor string
}