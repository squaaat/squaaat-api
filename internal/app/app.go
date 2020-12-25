package app

import (
	"github.com/rs/zerolog/log"

	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/db"
)

type Application struct {
	ServiceDB *db.Client
}

func New() *Application {
	a := &Application{}
	a.ServiceDB = CreateServiceDBClient()

	return a
}

func CreateServiceDBClient() *db.Client {
	client, err := db.New(config.ServiceDB, config.App)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	return client
}
