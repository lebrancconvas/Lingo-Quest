package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Displayname string `json:"displayname"`
}

func main() {
	app := fiber.New();
	
	// Middleware

	app.Use(cors.New());
	app.Use(logger.New());
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Lingo Quest API.");
	})

	var users []User;

	app.Post("/api/v1/users", func(c *fiber.Ctx) error {
		newUser := User{
			Username: "lebrancconvas",
			Password: "asidw9qw",
			Displayname: "Lebranc Convas",
		}

		users := append(users, newUser);
		return c.JSON(users); 
	})
	
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(users); 
	})

	app.Listen(":8009");
}