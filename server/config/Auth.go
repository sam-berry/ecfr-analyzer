package config

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
)

var TokenSecret = []byte(os.Getenv("ECFR_TOKEN_SECRET"))

var TokenHandler = jwtware.New(
	jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: TokenSecret},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		},
	},
)
