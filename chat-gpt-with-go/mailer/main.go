package main

import (
	"go.tienngay/mailer/database"
	"go.tienngay/mailer/router"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.MysqlConnectDB()
	database.MongoConnectDB()
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":2211"))
}
