package http

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	auth "github.com/jaysonmulwa/cuisine/internal/handler/auth"
	posts "github.com/jaysonmulwa/cuisine/internal/handler/posts"
	users "github.com/jaysonmulwa/cuisine/internal/handler/users"
)

func SetupRoutes() {

	fmt.Println("Hello, World!")

	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)
	app.Post("/logout", auth.Logout)
	app.Post("/refresh", auth.Refresh)

	app.Get("/users", users.GetUsers)
	app.Get("/users/:id", users.GetUser)
	app.Post("/users", users.CreateUser)
	app.Put("/users/:id", users.UpdateUser)
	app.Delete("/users/:id", users.DeleteUser)

	app.Get("/posts", posts.GetAll)
	app.Get("/posts/:id", posts.GetPost)
	app.Get("/posts/:user_id", posts.GetPostsByUser)
	app.Post("/posts", posts.CreatePost)
	app.Put("/posts/:id", posts.UpdatePost)
	app.Delete("/posts/:id", posts.DeletePost)

	app.Listen(":4000")

}
