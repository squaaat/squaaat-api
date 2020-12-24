package app

import (
	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/db"
)

func StartGORMInitialize() {
	dbClient := db.New(config.ServiceDB)
	dbClient.Initialize()
}
func StartGORMGenerate() {
	dbClient := db.New(config.ServiceDB)
