package cmd

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
	serverhttp "github.com/squaaat/squaaat-api/internal/server/http"
	"github.com/squaaat/squaaat-api/internal/server/http/v1/auth"

	"net"
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

		config.MustInit(env)
		app := app.New()
		http := serverhttp.New()

		http.Post("/api/v1/auth/login", auth.PostAuthLogin(app))
		http.Get("/swagger/*", swagger.Handler)

		host := net.JoinHostPort("0.0.0.0", config.ServerHTTP.Port)
		if err := http.Listen(host); err != nil {
			log.Fatal().Msgf("%v", err)
		}
	}

	return c
}
