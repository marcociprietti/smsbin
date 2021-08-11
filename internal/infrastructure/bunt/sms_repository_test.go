package bunt_test

import (
	"context"
	"github.com/cipma/smsbin/app/sms"
	"github.com/cipma/smsbin/infrastructure/bunt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSingleSms(t *testing.T) {
	t.Parallel()
	repo := bunt.NewBuntSmsRepository(bunt.ClientOptions{
		Path: "",
		Name: ":memory:",
	})

	smsUuid := uuid.NewString()
	smsWhen := time.Now()
	err := repo.Save(context.TODO(), sms.Sms{
		Uuid:         smsUuid,
		Text:         "test sms",
		From:         "+393331234567",
		To:           "+393339876543",
		When:         smsWhen,
		FromMetadata: sms.PhoneNumberMetadata{},
		ToMetadata:   sms.PhoneNumberMetadata{},
		TextMetadata: sms.TextMetadata{},
	})
	assert.NoError(t, err)

	singleSms, err := repo.GetSingleSms(context.TODO(), smsUuid)
	assert.NoError(t, err)
	assert.Equal(t, smsUuid, singleSms.Uuid)
	assert.Equal(t, "test sms", singleSms.Text)
	assert.Equal(t, "+393331234567", singleSms.From)
	assert.Equal(t, "+393339876543", singleSms.To)
	assert.Equal(t, smsWhen.Format(time.RFC3339Nano), singleSms.When.Format(time.RFC3339Nano))
}

func TestAllSmsPaginated(t *testing.T) {

}