package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	CookieExpiresAfter = time.Hour * 24
	ClientCookieName   = "CLIENTID"
)

func GetClientID(c *fiber.Ctx) string {
	return c.Cookies(ClientCookieName)
}

func IdentityMiddleware(c *fiber.Ctx) error {
	id := c.Cookies(ClientCookieName, "")
	if id == "" {
		id = uuid.NewString()
		c.Cookie(&fiber.Cookie{
			Name:    ClientCookieName,
			Value:   id,
			Expires: time.Now().Add(CookieExpiresAfter),
		})
	}
	return c.Next()
}
