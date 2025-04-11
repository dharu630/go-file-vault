package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const UploadDir = "./backend/uploads"

var encryptionKey = []byte("a very secret key!") // 16/24/32 bytes for AES-128/192/256
func encryptAndSaveFile(file io.Reader, filePath string) error {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}
	plainData, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	encrypted := gcm.Seal(nonce, nonce, plainData, nil)
	return os.WriteFile(filePath, encrypted, 0644)
}

func UploadFile(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "File not found in request"})
	}
	openedFile, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to open uploaded file"})
	}
	defer openedFile.Close()
	filePath := filepath.Join(UploadDir, file.Filename)
	err = encryptAndSaveFile(openedFile, filePath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to encrypt and save file"})
	}
	return c.JSON(fiber.Map{
		"message":  "Encrypted upload successful",
		"filename": file.Filename,
	})
}

func ListFiles(c *fiber.Ctx) error {
	files, err := os.ReadDir(UploadDir)
	if err != nil {
		return c.Status(500).SendString("Could not list files")
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
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return c.Status(500).SendString("Failed to read file")
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return c.Status(500).SendString("Failed to initialize cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return c.Status(500).SendString("Failed to initialize GCM")
	}
	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return c.Status(400).SendString("Corrupted encrypted file")
	}
	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return c.Status(500).SendString("Failed to decrypt file")
	}
	c.Response().Header.Set("Content-Disposition", "attachment; filename="+filename)
	c.Response().Header.Set("Content-Type", "application/octet-stream")
	return c.Send(plaintext)
}
