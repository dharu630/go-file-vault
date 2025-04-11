package main

import (
	"os"

	"github.com/dharmika_kanderi/secure-file-share/backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()

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
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
	app.Use("/files", jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	app.Listen(":3000")
}
