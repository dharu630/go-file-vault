package controllers

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const UploadDir = "./backend/uploads"

func UploadFile(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "http://localhost:5173") // Explicitly set the header

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "File not found in request"})
	}

	filePath := filepath.Join(UploadDir, file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to save file"})
	}

	return c.JSON(fiber.Map{
		"message":  "Uploaded successfully",
		"filename": file.Filename,
	})
}
func ListFiles(c *fiber.Ctx) error {
	files, err := os.ReadDir(UploadDir)
	if err != nil {
		return c.Status(500).SendString(" Could not list files")
	}

	fileList := []string{}
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	return c.JSON(fileList)
}

func DownloadFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join(UploadDir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).SendString("File not found")
	}

	return c.SendFile(filePath)
}
