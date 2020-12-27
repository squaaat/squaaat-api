package serverhttp

import (
	"fmt"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/squaaat/squaaat-api/docs"
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/server/http/v1/auth"
)

type HTTPApplication struct {
	App *app.Application
	HTTP *fiber.App
}


func New(a *app.Application) *HTTPApplication {
	httpApp := &HTTPApplication{}
	f := fiber.New()
	httpApp.App = a
	httpApp.HTTP = f

	httpApp.HTTP.Use(func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())
		fmt.Println(string(ctx.Body()))
		return ctx.Next()
	})

	httpApp.HTTP.Post("/api/v1/auth/login", auth.PostAuthLogin(a))

	httpApp.HTTP.Get("/swagger/*", swagger.Handler)

	return httpApp
}

