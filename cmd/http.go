package main

import (

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
)

func init() {
	config.MustInit()
}

func main() {
	app.StartHTTP()
}
