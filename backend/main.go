package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Displayname string `db:"displayname" json:"displayname"`
	Created_at string `db:"created_at" json:"created_at"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Displayname string `json:"displayname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sqlx.DB; 

func main() {
	var err error;
	db, err = sqlx.Open("mysql", "root:<SECRET_PASSWORD>@tcp(0.0.0.0:3306)/lingo_quest");
	if err != nil {
		fmt.Println()
		panic(err);
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	});
	
	// Middleware
	app.Use(cors.New());
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}));
	
	// CRUD
	// GET METHOD. 
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Lingo Quest API.");
	})

	v1 := app.Group("/api/v1", func (c *fiber.Ctx) error {
		c.Set("version", "v1");
		return c.Next(); 
	});

	v1.Post("/signup", signup);
	v1.Post("/login", login);
	v1.Get("/users", getUsers);

	app.Listen(":8009");
}

func signup(c *fiber.Ctx) error {
	request := SignupRequest{};
	err := c.BodyParser(&request);	
	if err != nil {
		return err;
	}

	if request.Username == "" || request.Password == "" || request.Displayname == "" {
		return fiber.ErrUnprocessableEntity; 
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 4);
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error());
	}

	query := "INSERT users (username, password, displayname) VALUES (?, ?, ?)";
	newGenPassword := string(password)[0:16];
	result, err := db.Exec(query, request.Username, newGenPassword, request.Displayname);
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error());
	}

	id, err := result.LastInsertId();
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error());
	}

	created_at := time.Now().Format(time.RFC850);

	user := User{
		ID: int(id),
		Username: request.Username,
		Password: newGenPassword,
		Displayname: request.Displayname,
		Created_at: created_at,
	}

	return c.Status(fiber.StatusCreated).JSON(user); 
}

func login(c *fiber.Ctx) error {
	return nil;
}

func getUsers(c *fiber.Ctx) error {
	return nil;
}