package command

import (
	"context"
	"github.com/cipma/smsbin/app/sms"
)

type SaveSmsCommand struct {
	Uuid string
	Text string
	From string
	To   string
}

type SaveSmsHandler struct {
	smsRepo     sms.Repository
	phoneParser sms.PhoneNumberParser
	textParser  sms.TextSegmentParser
}

func NewSaveSmsHandler(smsRepo sms.Repository, phoneParser sms.PhoneNumberParser, textParser sms.TextSegmentParser) SaveSmsHandler {
	if smsRepo == nil {
		panic("SMS repository is nil")
	}

	if phoneParser == nil {
		panic("PhoneNumberParser is nil")
	}

	return SaveSmsHandler{
		smsRepo:     smsRepo,
		phoneParser: phoneParser,
		textParser:  textParser,
	}
}

func (h SaveSmsHandler) Handle(ctx context.Context, cmd SaveSmsCommand) error {
	newSms, err := sms.NewSms(cmd.Uuid, cmd.Text, cmd.From, cmd.To, h.phoneParser, h.textParser)
	if err != nil {
		return err
	}

	err = h.smsRepo.Save(ctx, newSms)
	if err != nil {
		return err
	}

	return nil
}
