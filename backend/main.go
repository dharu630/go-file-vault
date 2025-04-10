package main

import (
	"github.com/dharmika_kanderi/secure-file-share/backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/test-cors", func(c *fiber.Ctx) error {
		return c.SendString("CORS Test Successful")
	})

	app.Post("/upload", controllers.UploadFile)
	app.Get("/files", controllers.ListFiles)
	app.Get("/download/:filename", controllers.DownloadFile)

	app.Listen(":3000")
}
