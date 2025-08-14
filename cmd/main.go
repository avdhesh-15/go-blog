package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avdhesh-15/go-blog-backend/config"
	"github.com/avdhesh-15/go-blog-backend/controllers"
	"github.com/avdhesh-15/go-blog-backend/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error:", err)
	}
	PORT := os.Getenv("PORT")

	config.DbConfig()
	fmt.Println("✅ Connected to Postgres")
	defer config.DbPool.Close()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Post("api/signup", controllers.SignUp)
	app.Post("api/signin", controllers.SignIn)

	app.Post("api/admin/signup", controllers.AdminSignUp)

	app.Get("api/blogs", middlewares.JwtVerify(), controllers.GetAllBlogs)
	app.Get("api/blogs/blog/:id", middlewares.JwtVerify(), controllers.GetBlog)
	app.Post("api/blogs/blog/create", middlewares.JwtVerify(), middlewares.RoleRequired("admin"), controllers.CreateBlog)
	app.Delete("api/blogs/blog/:id", middlewares.JwtVerify(), middlewares.RoleRequired("admin"), controllers.DeleteBLog)

	app.Get("api/profile", middlewares.JwtVerify(), func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Welcome, User"})
	})
	app.Get("api/admin/profile", middlewares.JwtVerify(), middlewares.RoleRequired("admin"), func(c *fiber.Ctx) error {
		name := c.Locals("user_name").(string)
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": fmt.Sprintf("Hello there admin, %s", name)})
	})
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(":" + PORT))

}
