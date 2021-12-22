package http

import (
	"fmt"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	auth "github.com/jaysonmulwa/cuisine/internal/handler/auth"
	comments "github.com/jaysonmulwa/cuisine/internal/handler/comments"
	likes "github.com/jaysonmulwa/cuisine/internal/handler/likes"
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

	app.Get("/comments/:id", comments.GetComment)
	app.Get("/comments/:post_id", comments.GetCommentsByPost)
	app.Get("/comments/:post_id/:user_id", comments.GetCommentsByPostForUser)
	app.Post("/comments", comments.CreateComment)
	app.Put("/comments/:id", comments.UpdateComment)
	app.Delete("/comments/:id", comments.DeleteComment)

	app.Get("/likes/:post_id", likes.GetLikesByPost)
	app.Get("/likes/:post_id/:user_id", likes.GetLikesByPostForUser)
	app.Post("/likes", likes.CreateLike)
	app.Delete("/likes/:id", likes.DeleteLike)

	app.Listen(":4000")

}
