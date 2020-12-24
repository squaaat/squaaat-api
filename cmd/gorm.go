package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
)

func newGormCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "gorm",
		Short: "squaaat-api cli gorm scripts",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newGormGenerateCommand())
	c.AddCommand(newGormInitializeCommand())

	return c
}

func newGormInitializeCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "init",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		runGormInitialize(&Options{
			Env: env,
		})
	}

	return c
}

func newGormGenerateCommand() *cobra.Command {
	c := &cobra.Command{
		Use: "generate",
		Short: "re-create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run http server")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		runGormGenerator(&Options{
			Env: env,
		})
	}

	return c
}

func runGormInitialize(o *Options) {
	config.MustInit(o.Env)
	app.StartGORMInitialize()
}

func runGormGenerator(o *Options) {
	config.MustInit(o.Env)
	app.StartGORMGenerate()
}