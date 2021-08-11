package app

import (
	"github.com/cipma/smsbin/app/command"
	"github.com/cipma/smsbin/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SaveSmsHandler command.SaveSmsHandler
}

type Queries struct {
	FindSmsHandler query.FindSmsHandler
	AllSmsHandler  query.AllSmsHandler
}
