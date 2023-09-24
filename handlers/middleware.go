package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const CookieExpiresAfter = time.Hour * 24

func IdentityMiddleware(c *fiber.Ctx) error {
    id := c.Cookies("uuid", "")
    if id == "" {
        id = uuid.NewString()
        c.Cookie(&fiber.Cookie{
            Name: "uuid",
            Value: id,
            Expires: time.Now().Add(CookieExpiresAfter),
        })
    } 
    return c.Next()
}
