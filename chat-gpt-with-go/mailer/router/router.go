package router

import (
	"go.tienngay/mailer/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	mailer := app.Group("/mailer", logger.New())
	mailer.Get("/send-email", handler.SendEmail)
}
