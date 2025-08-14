package controllers

import (
	"context"
	"fmt"
	"strings"

	"os"
	"time"

	"github.com/avdhesh-15/go-blog-backend/config"
	"github.com/avdhesh-15/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userData); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Incorrect content"})
	}

	if strings.TrimSpace(userData.Username) == "" ||
		strings.TrimSpace(userData.Email) == "" ||
		strings.TrimSpace(userData.Password) == "" {
		return c.Status(400).JSON(fiber.Map{"message": "All fields are required"})
	}

	userEmail := userData.Email

	err := config.DbPool.QueryRow(context.Background(),
		"SELECT id from users WHERE email = $1",
		userEmail,
	).Scan()

	if err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "user already exist"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	_, err = config.DbPool.Exec(
		context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		userData.Username, userData.Email, hashedPassword,
	)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "user created successfully"})
}

func SignIn(c *fiber.Ctx) error {
	var userData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Incorrect content"})
	}

	if strings.TrimSpace(userData.Email) == "" ||
		strings.TrimSpace(userData.Password) == "" {
		return c.Status(400).JSON(fiber.Map{"message": "All fields are required"})
	}

	user := models.User{}

	err := config.DbPool.QueryRow(context.Background(),
		"SELECT id, name, password, role FROM users WHERE email =$1",
		userData.Email).Scan(
		&user.Id,
		&user.UserName,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid password"})
	}

	err = godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error:", err)
	}

	claims := jwt.MapClaims{
		"user_id":   user.Id,
		"user_name": user.UserName,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	return c.JSON(fiber.Map{"token": secret})

}

// Admin Auth Portals

func AdminSignUp(c *fiber.Ctx) error {
	var userData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userData); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Incorrect content"})
	}
	if strings.TrimSpace(userData.Username) == "" ||
		strings.TrimSpace(userData.Email) == "" ||
		strings.TrimSpace(userData.Password) == "" {
		return c.Status(400).JSON(fiber.Map{"message": "All fields are required"})
	}

	userEmail := userData.Email
	err := config.DbPool.QueryRow(context.Background(),
		"SELECT id from users WHERE email = $1",
		userEmail,
	).Scan()

	if err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "user already exist"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	_, err = config.DbPool.Exec(
		context.Background(),
		"INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)",
		userData.Username, userData.Email, hashedPassword, "admin",
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "user created successfully"})
}
