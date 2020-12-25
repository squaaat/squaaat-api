package app

import (
	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/db"
)

func StartGORMInit() {
	db.Initialize(config.ServiceDB)
}

func StartGORMClean() {
	dbClient := db.New(config.ServiceDB, config.App)
	dbClient.Clean()
}
