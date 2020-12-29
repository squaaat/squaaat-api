package serverhttp

import (
	"fmt"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/server/http/v1/auth"
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

type HTTPApplication struct {
	App  *app.Application
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
