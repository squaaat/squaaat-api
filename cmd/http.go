package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
)

func newHTTPCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "http",
		Short: "about squaaat-api http server",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newHTTPStartCommand())

	return c
}

func newHTTPStartCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "start",
		Short: "run http application",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		runHTTPServer(&Options{
			Env: env,
		})
	}

	return c
}

type Options struct {
	Env string
}

func runHTTPServer(o *Options) {
	config.MustInit(o.Env)
	app.StartHTTP()
}
