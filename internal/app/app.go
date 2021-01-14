package app

import (
	"github.com/rs/zerolog/log"

	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/db"
)

type Application struct {
	Config    *config.Config
	ServiceDB *db.Client
}

func New(cfg *config.Config) *Application {
	a := &Application{
		Config: cfg,
	}
	a.ServiceDB = CreateServiceDBClient(a.Config)

	return a
}

func CreateServiceDBClient(cfg *config.Config) *db.Client {
	client, err := db.New(db.ParseConfig(cfg))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	return client
}
