package gorm

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/db"
	"github.com/squaaat/squaaat-api/migrations"
	"os"
	"strconv"
	"time"
)

const (
	ArgEnv        = "environment"
	ArgEnvShort   = "e"
	ArgEnvDefault = "alpha"

	ArgVersion      = "version"
	ArgVersionShort = "v"
)

// -------- gorm ------------------------
func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "gorm",
		Short: "squaaat-api cli gorm scripts",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(newGormClean())
	c.AddCommand(newGormCreate())
	c.AddCommand(newGormReCreate())
	c.AddCommand(newGormMigrate())

	return c
}

func newGormReCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "re-create",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		sqcicd := os.Getenv("SQ_CICD")
		cicd, _ := strconv.ParseBool(sqcicd)
		cfg := config.MustInit(env, cicd)
		a := app.New(cfg)
		a.ServiceDB.Clean()

		err = db.Initialize(cfg.ServiceDB)

		m := migrations.New(a)
		m.Sync()

		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}

	return c
}

func newGormCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "create schema for develop",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		sqcicd := os.Getenv("SQ_CICD")
		cicd, _ := strconv.ParseBool(sqcicd)
		cfg := config.MustInit(env, cicd)

		err = db.Initialize(cfg.ServiceDB)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		a := app.New(cfg)
		m := migrations.New(a)
		m.Sync()

		if err != nil {
			log.Fatal().Err(err).Send()
		}
	}

	return c
}

func newGormClean() *cobra.Command {
	c := &cobra.Command{
		Use:   "clean",
		Short: "remove schema(db)",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		sqcicd := os.Getenv("SQ_CICD")
		cicd, _ := strconv.ParseBool(sqcicd)
		cfg := config.MustInit(env, cicd)
		a := app.New(cfg)
		a.ServiceDB.Clean()
	}

	return c
}

// -------- gorm migrate ------------------------
func newGormMigrate() *cobra.Command {
	c := &cobra.Command{
		Use:   "migrate",
		Short: "it's about gorm migrate",
	}
	c.Run = func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	}

	c.AddCommand(newGormMigrateSync())
	c.AddCommand(newGormMigrateCreate())

	return c
}

func newGormMigrateSync() *cobra.Command {
	c := &cobra.Command{
		Use:   "sync",
		Short: "sync migrations code",
	}

	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		sqcicd := os.Getenv("SQ_CICD")
		cicd, _ := strconv.ParseBool(sqcicd)
		cfg := config.MustInit(env, cicd)
		a := app.New(cfg)
		m := migrations.New(a)
		m.Sync()
	}

	return c
}

func newGormMigrateCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "sync migrations code",
	}
	c.Flags().StringP(ArgEnv, ArgEnvShort, ArgEnvDefault, "set environment to run gorm cli")
	c.Flags().StringP(ArgVersion, ArgVersionShort, time.Now().Format("200601021504"), "set version to create migration")
	c.Run = func(cmd *cobra.Command, _ []string) {
		env, err := cmd.Flags().GetString(ArgEnv)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		version, err := cmd.Flags().GetString(ArgVersion)
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		sqcicd := os.Getenv("SQ_CICD")
		cicd, _ := strconv.ParseBool(sqcicd)
		cfg := config.MustInit(env, cicd)
		a := app.New(cfg)
		m := migrations.New(a)
		m.Create(version)
	}

	return c
}
