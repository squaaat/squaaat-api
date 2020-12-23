package main

import (
	"github.com/getsentry/sentry-go"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/mornitor"
)

func init() {
	config.MustInit()
	mornitor.MustInit(config.App, config.Sentry)
}

func main() {
	app.StartHTTP()
	sentry.CaptureMessage("wow!")
	sentry.CaptureMessage("wow!")
	defer sentry.Flush(config.Sentry.FlushTimeout)
}
