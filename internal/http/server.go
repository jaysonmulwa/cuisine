package http

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	posts "github.com/jaysonmulwa/cuisine/internal/handler/posts"
)

func SetupRoutes() {

	fmt.Println("Hello, World!")

	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/posts", posts.GetAll)
	app.Get("/posts/:id", posts.GetPost)
	app.Get("/posts/:user_id", posts.GetPostsByUser)
	app.Post("/posts", posts.CreatePost)
	app.Put("/posts/:id", posts.UpdatePost)
	app.Delete("/posts/:id", posts.DeletePost)

	app.Listen(":4000")

}
