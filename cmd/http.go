package main

import (
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
)

func main() {
	config.MustInit()
	app.StartHTTP()
}
