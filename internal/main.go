package main

import (
	"context"
	"github.com/cipma/smsbin/infrastructure"
	"github.com/cipma/smsbin/infrastructure/http"
)

func main() {
	ctx := context.Background()
	application := infrastructure.NewApplication(ctx)

	http.RunHttpServer(application)
}
