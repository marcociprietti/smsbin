package phonenumber

import (
	"github.com/cipma/smsbin/app/sms"
	"github.com/nyaruka/phonenumbers"
)

type Parser struct {
	defaultRegion string
}

func NewPhoneNumberParser(options ParserOptions) Parser {
	return Parser{
		defaultRegion: options.DefaultRegion,
	}
}

func (p Parser) Parse(number string, allowAlpha bool) sms.PhoneNumberMetadata {
	if allowAlpha && p.isAlphanumericSender(number) {
		return sms.PhoneNumberMetadata{
			IsValid:                true,
			IsAlphanumeric:         true,
			InternationalFormatted: number,
		}
	}

	metadata, err := phonenumbers.Parse(number, p.defaultRegion)
	if err != nil || !phonenumbers.IsValidNumber(metadata) {
		return sms.PhoneNumberMetadata{}
	}

	return sms.PhoneNumberMetadata{
		CountryCode:            *metadata.CountryCode,
		NationalNumber:         *metadata.NationalNumber,
		RegionCode:             phonenumbers.GetRegionCodeForCountryCode(int(*metadata.CountryCode)),
		IsValid:                phonenumbers.IsValidNumber(metadata),
		InternationalFormatted: phonenumbers.Format(metadata, phonenumbers.INTERNATIONAL),
	}
}

//isAlphanumericSender check if number is a valid alphanumeric sender ID
func (p Parser) isAlphanumericSender(number string) bool {
	if len(number) == 0 || len(number) > 11 {
		return false
	}

	var hasLetter bool
	for _, r := range number {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			hasLetter = true
			continue
		}

		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') && r != ' ' {
			return false
		}
	}

	return hasLetter
}
