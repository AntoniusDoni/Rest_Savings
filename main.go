package main

import (
	"github.com/Rest_Savings/apps/services"
	_ "github.com/Rest_Savings/docs"
	"github.com/Rest_Savings/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Fiber Swagger Saving API
// @version 2.0
// @description This is a Rest Saving server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @schemes http

func main() {
	apps := fiber.New(fiber.Config{
		// Views: engine,
	})
	apps.Use(cors.New())
	apps.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	apps.Use(recover.New())
	services := services.New()
	services.Db.GetInstance()
	defer services.Db.Conn.Close()
	apps.Get("/swagger/*", swagger.HandlerDefault)

	route.Routes(apps, services)

	apps.Listen(":3000")
}
