package infrastructure

import (
	"context"
	"github.com/cipma/smsbin/app"
	"github.com/cipma/smsbin/app/command"
	"github.com/cipma/smsbin/app/query"
	"github.com/cipma/smsbin/infrastructure/config"
	"github.com/cipma/smsbin/infrastructure/phonenumber"
	"github.com/cipma/smsbin/infrastructure/segment"
	"github.com/cipma/smsbin/infrastructure/storm"
)

func NewApplication(_ context.Context) app.Application {
	configuration := config.GetConfig()

	smsRepository := storm.NewSmsRepository(storm.Options{
		Path: configuration.GetString("database.path"),
		Name: configuration.GetString("database.name"),
	})
	phoneParser := phonenumber.NewPhoneNumberParser(phonenumber.ParserOptions{
		DefaultRegion: configuration.GetString("default_region"),
	})
	textParser := segment.Parser{}

	return app.Application{
		Commands: app.Commands{
			SaveSmsHandler: command.NewSaveSmsHandler(smsRepository, phoneParser, textParser),
		},
		Queries: app.Queries{
			FindSmsHandler: query.NewFindSmsHandler(smsRepository),
			AllSmsHandler:  query.NewAllSmsHandler(smsRepository),
		},
	}
}
