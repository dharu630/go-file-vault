package controllers

import (
	"os"
	"time"

	"github.com/dharmika_kanderi/secure-file-share/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)

func Signup(c *fiber.Ctx) error {
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	users[u.Username] = string(hash)

	return c.JSON(fiber.Map{"message": "Signup successful"})
}
func Login(c *fiber.Ctx) error {
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}
	stored, ok := users[u.Username]
	if !ok || bcrypt.CompareHashAndPassword([]byte(stored), []byte(u.Password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
