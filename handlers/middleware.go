package handlers

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/auth"
)

const (
	CookieExpiresAfter = time.Hour * 24
	ClientCookieName   = "USERID"
)

func GetUserID(c *fiber.Ctx) string {
	return c.Cookies(ClientCookieName)
}

func IdentityMiddlewareDevice(c *fiber.Ctx) error {
	ip := c.IP()
	identity, err := auth.MakeIdentity(ip)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return c.Render("auth", nil)
		}
		return err
	}

	log.Printf("Identity: %v", identity)
	c.Locals("Identity", identity)

	return c.Next()
}

func IdentityMiddlewareCoockie(c *fiber.Ctx) error {
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
