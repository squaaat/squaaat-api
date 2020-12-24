package cmd

import (
	"github.com/rs/zerolog/log"
)

const (
	ArgEnv      = "environment"
	ArgEnvShort = "e"
	ArgEnvDefault = "alpha"
)


func Start() {
	c := newCliCmd()
	c.AddCommand(newHTTPCommand())
	c.AddCommand(newGormCommand())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
