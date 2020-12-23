package mornitor

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog/log"

	"github.com/squaaat/squaaat-api/internal/config"
)

func MustInit(appcfg *config.AppConfig, cfg *config.SentryConfig) {
	if !cfg.Enabled {
		return
	}

	fmt.Println("hello")
	err := sentry.Init(sentry.ClientOptions{
		Release:          config.Version,
		ServerName:       appcfg.AppName,
		Environment:      appcfg.Env,
		Dsn:              cfg.DSN,
		Debug:            appcfg.Debug,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	// Flush buffered events before the program terminates.

	sentry.CaptureMessage("It works!")
}
