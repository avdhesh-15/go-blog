package controllers

import (
	"context"
	"fmt"

	"github.com/avdhesh-15/go-blog-backend/config"
	"github.com/avdhesh-15/go-blog-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetAllBlogs(c *fiber.Ctx) error {
	rows, err := config.DbPool.Query(context.Background(), "SELECT id, title, content, author_id FROM posts")
	if err != nil {
		c.Status(400).JSON(fiber.Map{"error": "Cannot fetch Blogs"})
	}
	blogs := []models.Post{}

	for rows.Next() {
		var blog models.Post

		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.AuthorId); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan blog"})
		}
		blogs = append(blogs, blog)
	}

	return c.Status(fiber.StatusAccepted).JSON(blogs)
}

func GetBlog(c *fiber.Ctx) error {
	blogId := c.Params("id")

	blog := config.DbPool.QueryRow(context.Background(),
		"SELECT * FROM posts WHERE id = $1",
		blogId,
	)

	if blog == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Blog not found"})
	}

	return c.Status(fiber.StatusAccepted).JSON(blog)
}

func CreateBlog(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	var authorId int
	if uid, ok := claims["user_id"].(float64); ok {
		authorId = int(uid)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized author"})
	}

	blog := models.Post{}

	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ivalid Request"})
	}

	if blog.Title == "" || blog.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Title and Content Required"})
	}

	_, err := config.DbPool.Exec(context.Background(), "INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)", blog.Title, blog.Content, authorId)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create blog"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Blog created successfully",
	})
}

func DeleteBLog(c *fiber.Ctx) error {
	blogId := c.Params("id")

	_, err := config.DbPool.Exec(context.Background(), "DELETE FROM posts WHERE id = $1", blogId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot delete the blog",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Blog deleted successfully",
	})

}
