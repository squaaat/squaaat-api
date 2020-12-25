package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
)

func newGormCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "gorm",
		Short: "squaaat-api cli gorm scripts",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newGormClean())
	c.AddCommand(newGormInit())

	return c
}

func newGormInit() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		runGormInit(&Options{
			Env: env,
		})
	}

	return c
}

func newGormClean() *cobra.Command {
	c := &cobra.Command{
		Use:   "clean",
		Short: "remove schema(db)",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		runGormClean(&Options{
			Env: env,
		})
	}

	return c
}

func runGormInit(o *Options) {
	config.MustInit(o.Env)
	app.StartGORMInit()
}

func runGormClean(o *Options) {
	config.MustInit(o.Env)
	app.StartGORMClean()
}
