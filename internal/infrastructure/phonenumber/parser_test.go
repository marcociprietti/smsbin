package phonenumber_test

import (
	"fmt"
	"github.com/cipma/smsbin/infrastructure/phonenumber"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	number              string
	defaultRegion       string
	allowAlpha          bool
	countryCode         int32
	nationalNumber      uint64
	isValid             bool
	isAlphanumeric      bool
	regionCode          string
	internationalFormat string
}{
	{
		"3331234567",
		"IT",
		false,
		int32(39),
		uint64(3331234567),
		true,
		false,
		"IT",
		"+39 333 123 4567",
	},
	{
		"+393331234567",
		"US",
		false,
		int32(39),
		uint64(3331234567),
		true,
		false,
		"IT",
		"+39 333 123 4567",
	},
	{
		"+12025886500",
		"US",
		false,
		int32(1),
		uint64(2025886500),
		true,
		false,
		"US",
		"+1 202-588-6500",
	},
	{
		"Alpha 123",
		"US",
		true,
		0,
		0,
		true,
		true,
		"",
		"Alpha 123",
	},
	{
		"Alpha 123",
		"US",
		false,
		0,
		0,
		false,
		false,
		"",
		"",
	},
	{
		"",
		"US",
		false,
		0,
		0,
		false,
		false,
		"",
		"",
	},
}

func TestParse(t *testing.T) {
	t.Parallel()

	for i, td := range testData {
		td := td
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			t.Parallel()

			parser := phonenumber.NewPhoneNumberParser(phonenumber.ParserOptions{
				DefaultRegion: td.defaultRegion,
			})
			metadata := parser.Parse(td.number, td.allowAlpha)

			assert.Equal(t, td.countryCode, metadata.CountryCode)
			assert.Equal(t, td.nationalNumber, metadata.NationalNumber)
			assert.Equal(t, td.isValid, metadata.IsValid)
			assert.Equal(t, td.isAlphanumeric, metadata.IsAlphanumeric)
			assert.Equal(t, td.regionCode, metadata.RegionCode)
			assert.Equal(t, td.internationalFormat, metadata.InternationalFormatted)
		})
	}
}
