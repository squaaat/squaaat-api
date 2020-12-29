package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
	"github.com/squaaat/squaaat-api/internal/model"
	"github.com/squaaat/squaaat-api/pkg/rsautil"
)

// swagger:route POST /auth/login webview client logins
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       default: genericError
//       204: someResponse
//       400: validationError
//       500: serverError

type LoginRequest struct {
	DeviceToken string `json:"device_token"`
}

func ParseAuthRequest(ctx *fiber.Ctx) (*LoginRequest, error) {
	req := &LoginRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return nil, err
	}

	return req, nil
}

func PostAuthLogin(app *app.Application) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req, err := ParseAuthRequest(ctx)
		if err != nil {
			log.Error().Err(err).Send()
			return ctx.
				Status(fiber.StatusBadRequest).
				SendString(err.Error())
		}

		var userSessionsTokens []model.UserDevice
		ud := &model.UserDevice{DeviceToken: req.DeviceToken}
		app.ServiceDB.DB.Find(&userSessionsTokens, ud)
		if len(userSessionsTokens) == 0 {
			tx := app.ServiceDB.DB.Begin()
			u := &model.User{}

			r := tx.Create(&u)
			if r.Error != nil {
				tx.Rollback()
				log.Error().Err(r.Error).Send()
				return ctx.
					Status(fiber.StatusInternalServerError).
					SendString(r.Error.Error())
			}

			ud.UserID = u.ID
			ud.UpdatedBy = u.ID
			ud.CreatedBy = u.ID

			r = tx.Create(ud)
			if r.Error != nil {
				tx.Rollback()
				log.Error().Err(r.Error).Send()
				return ctx.
					Status(fiber.StatusInternalServerError).
					SendString(r.Error.Error())
			}

			tx.Commit() // Commit user1
		} else if len(userSessionsTokens) == 1 {
			ud = &userSessionsTokens[0]
		} else {
			// FIXME: 서버로 반드시 알림을 만들어서 요청을 보내도록, device_token이 중복되는 일은 없어야함.
		}

		// REFERME: jwt, jwk https://docs.aws.amazon.com/ko_kr/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
		t := jwt.New()
		t.Set(jwt.SubjectKey, `https://squaaat.com/api/v1`)
		t.Set(jwt.AudienceKey, `SQUAAAT Users`)
		t.Set(jwt.IssuedAtKey, time.Now().Unix())
		t.Set(jwt.ExpirationKey, time.Now().Add(3*time.Minute).Unix())
		t.Set("user_id", 1485)

		privateKey, err := rsautil.ParseRsaPrivateKeyFromPemStr(config.App.RSAPrivatePem)
		if err != nil {
			log.Error().Err(err).Send()
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}

		payload, err := jwt.Sign(t, jwa.RS256, privateKey)
		if err != nil {
			log.Error().Err(err).Send()
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}
		ctx.Set("X-Auth-Token", string(payload))
		return ctx.Status(fiber.StatusNoContent).Send([]byte(""))
	}
}
