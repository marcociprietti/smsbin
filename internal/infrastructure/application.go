package infrastructure

import (
	"context"
	"github.com/cipma/smsbin/app"
	"github.com/cipma/smsbin/app/command"
	"github.com/cipma/smsbin/app/query"
	"github.com/cipma/smsbin/infrastructure/bunt"
	"github.com/cipma/smsbin/infrastructure/config"
	"github.com/cipma/smsbin/infrastructure/phonenumber"
	"github.com/cipma/smsbin/infrastructure/segment"
)

func NewApplication(_ context.Context) app.Application {
	configuration := config.GetConfig()

	smsRepository := bunt.NewBuntSmsRepository(bunt.ClientOptions{
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
