package news

import (
	"github.com/gofiber/fiber/v2"
)

// News HTTP Handlers interface
type Handlers interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetNews(c *fiber.Ctx) error
}
