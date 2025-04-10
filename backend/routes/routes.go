package routes

import (
	"github.com/dharmika_kanderi/secure-file-share/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/upload", controllers.UploadFile)
	app.Get("/files", controllers.ListFiles)
	app.Get("/download/:filename", controllers.DownloadFile)
}
