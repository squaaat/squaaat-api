package serverhttp

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/service/auth"
)

// SQUAAAT application http api server that squaaat-api
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /api/v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: drakejin<dydwls121200@gmail.com> https://github.com/drakejin
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

func New(a *app.Application) *fiber.App {
	f := fiber.New()

	f.Use(func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Path())
		fmt.Println(string(ctx.Body()))
		return ctx.Next()
	})

	f.Post("/api/v1/auth/login", auth.PostAuthLogin(a))
	f.Get("/swagger/*", swagger.Handler)

	return f
}
