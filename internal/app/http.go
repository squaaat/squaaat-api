package app

import "github.com/gofiber/fiber/v2"

type Application struct {
	HTTP fiber.App
}

func StartHTTP() *Application {
	return &Application{}
}
