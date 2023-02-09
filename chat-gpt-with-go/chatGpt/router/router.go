package router

import (
	"go.tienngay/chatGpt/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	gpt := app.Group("/openai-gpt", logger.New())
	gpt.Post("/chat", handler.Chat)
}
