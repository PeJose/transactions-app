package middleware

import (
	"cashflow/pkg/config"
	"cashflow/pkg/models"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.SigningKey)},
		ErrorHandler: jwtError,
		TokenLookup:  "cookie:jwt-token",
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user")
			if user != nil {
				claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
				c.Locals("userClaims", claims)
			}
			return c.Next()
		},
	})
}

func GetUser(c *fiber.Ctx) jwt.MapClaims {
	userClaims := c.Locals("userClaims")
	if userClaims != nil {
		return userClaims.(jwt.MapClaims)
	}
	return nil
}

func GenerateToken(u *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iban"] = u.Iban

	signed, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	return signed, nil
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": err.Error()})
}
