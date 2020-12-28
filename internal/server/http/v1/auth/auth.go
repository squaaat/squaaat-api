package auth

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/model"
	"github.com/squaaat/squaaat-api/pkg/rsautil"
)


type LoginRequest struct {
	DeviceToken string `json:"device_token"`
}

type LoginResponse struct {
	PublicPem string `json:"public_pem"`
	JWTToken  string `json:"jwt_token"`
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
			return ctx.
				Status(fiber.StatusBadRequest).
				SendString(err.Error())
		}

		var userSessionsTokens []model.UserSessionToken
		ust := &model.UserSessionToken{DeviceToken: req.DeviceToken}
		app.ServiceDB.DB.Find(&userSessionsTokens, ust)
		if len(userSessionsTokens) == 0 {
			tx := app.ServiceDB.DB.Begin()
			u := &model.User{}

			r := tx.Create(&u)
			if r.Error != nil {
				tx.Rollback()
				return ctx.
					Status(fiber.StatusInternalServerError).
					SendString(r.Error.Error())
			}

			ust.UserID = u.ID
			ust.UpdatedBy = u.ID
			ust.CreatedBy = u.ID

			privateKey, _ := rsautil.GenerateRsaKeyPair(2048)
			ust.RSAPem = rsautil.ExportRsaPrivateKeyAsPemStr(privateKey)

			r = tx.Create(ust)
			if r.Error != nil {
				tx.Rollback()
				return ctx.
					Status(fiber.StatusInternalServerError).
					SendString(r.Error.Error())
			}

			tx.Commit() // Commit user1
		} else if len(userSessionsTokens) == 1 {
			ust = &userSessionsTokens[0]
		} else {
			// FIXME: 서버로 반드시 알림을 만들어서 요청을 보내도록, device_token이 중복되는 일은 없어야함.
		}

		privateKey, err := rsautil.ParseRsaPrivateKeyFromPemStr(ust.RSAPem)
		if err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}

		pub, err := rsautil.ExportRsaPublicKeyAsPemStr(&privateKey.PublicKey)
		if err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}

		// REFERME: jwt, jwk https://docs.aws.amazon.com/ko_kr/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
		const aLongLongTimeAgo = 233431200
		t := jwt.New()
		t.Set(jwt.SubjectKey, `https://squaaat.com/api/v1`)
		t.Set(jwt.AudienceKey, `SQUAAAT Users`)
		t.Set(jwt.IssuedAtKey, time.Now().Unix())
		t.Set(jwt.ExpirationKey, time.Now().Add(3*time.Minute).Unix())
		t.Set("user_id", 1485)

		payload, err := jwt.Sign(t, jwa.RS256, privateKey)
		if err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}
		res := &LoginResponse{
			JWTToken:  string(payload),
			PublicPem: pub,
		}

		if b, err := json.Marshal(res); err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		} else {
			return ctx.Send(b)
		}
	}
}
