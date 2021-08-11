package sms

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type noopPhoneNumberParser struct{}

func (n noopPhoneNumberParser) Parse(_ string, _ bool) PhoneNumberMetadata {
	return PhoneNumberMetadata{}
}

type noopTextSegmentParser struct{}

func (n noopTextSegmentParser) ExtractSegments(_ string) TextMetadata {
	return TextMetadata{}
}

func TestNewSms(t *testing.T) {
	t.Parallel()

	text := "Sms text"
	from := "3339876543"
	to := "3331234567"

	sms, err := NewSms(uuid.NewString(), text, from, to, noopPhoneNumberParser{}, noopTextSegmentParser{})

	assert.Nil(t, err)
	assert.Equal(t, text, sms.Text)
	assert.Equal(t, from, sms.From)
	assert.Equal(t, to, sms.To)
}

var testData = []struct {
	uuid string
	from string
	to   string
}{
	{uuid.NewString(), "", "3331234567"},
	{uuid.NewString(), "33398765432", ""},
	{"", "33398765432", "3331234567"},
}

func TestNewSms_invalid(t *testing.T) {
	t.Parallel()

	for _, td := range testData {
		td := td

		t.Run("", func(t *testing.T) {
			t.Parallel()

			_, err := NewSms(td.uuid, "Hello World!", td.from, td.to, noopPhoneNumberParser{}, noopTextSegmentParser{})
			assert.Error(t, err)
		})
	}
}
