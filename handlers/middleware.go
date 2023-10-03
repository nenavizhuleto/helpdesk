package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/auth"
	"application/models"
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
		if errors.Is(err, models.ErrUnsupportedDevice) {
			return c.Render("screens/unsupported-device-error", fiber.Map{}, "layout/main")
		}
		if errors.Is(err, models.ErrUserNotFound) {
			return c.Render("screens/authorization", fiber.Map{}, "layout/main")
		}
		return err
	}

	c.Locals("Identity", identity)

	return c.Next()
}

// TODO: After first identity we need to cache it to improve performance
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
