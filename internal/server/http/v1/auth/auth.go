package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/model"
	"gorm.io/gorm"
)


type Request struct {
	DeviceToken string `json:"device_token"`
}

func ParseAuthRequest(ctx *fiber.Ctx) (*Request, error){
	req := &Request{}
	if err := ctx.BodyParser(req); err != nil {
		return nil, err
	}

	return req, nil
}

func PostAuthLogin(app *app.Application) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req, err := ParseAuthRequest(ctx)
		if err != nil {
			return ctx.
				Status(fiber.StatusBadRequest).
				SendString(err.Error())
		}
		var userDevices []model.UserDevice
		app.ServiceDB.DB.Find(&userDevices, &model.UserDevice{DeviceToken: req.DeviceToken})
		if len(userDevices) == 0 {
			u := &model.User{}
			tx := app.ServiceDB.DB.Create(u)
			tx.Transaction(func (tc *gorm.DB) error {
				r := tc.Create(&model.UserDevice{
					UserID: u.ID,
					DeviceToken: req.DeviceToken,
				})
				if r.Error != nil {
					tc.Rollback()
				}
				tc.Commit()
				return r.Error
			})
		}

		return ctx.SendString("HelloWorld")
	}
}