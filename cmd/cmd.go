package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/squaaat-api/cmd/gorm"
)

const (
	ArgEnv        = "environment"
	ArgEnvShort   = "e"
	ArgEnvDefault = "alpha"
)

func newCliCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sq",
		Short: "squaaat-api application",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
}

func Start() {
	c := newCliCmd()
	c.AddCommand(newHTTPCommand())
	c.AddCommand(gorm.New())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
