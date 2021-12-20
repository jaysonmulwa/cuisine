package http

import (
	"context"
	"fmt"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID            int    `json:"id,omitempty" bson:"_id,omitempty"`
	Recipe_name   string `json:"recipe_name"`
	Ingredients   string `json:"ingredients"`
	Instructions  string `json:"instructions"`
	Origin        string `json:"origin"`
	Recipe_image  string `json:"recipe_image"`
	Register_date string `json:"register_date"`
	Post_type     string `json:"post_type"`
}

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	mg MongoInstance
)

// Connect configures the MongoDB client and initializes the database connection.
func Connect() error {

	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	mongoURI, ok := viper.Get("mongoURI").(string)
	dbName, ok := viper.Get("dbName").(string)

	fmt.Println(mongoURI)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func SetupRoutes() {

	fmt.Println("Hello, World!")

	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/posts", getAll)
	app.Get("/posts/:id", getPost)
	app.Post("/posts", createPost)
	app.Put("/posts/:id", updatePost)
	app.Delete("/posts/:id", deletePost)

	app.Listen(":4000")

}

func getAll(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("recipes").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var orders []Post = make([]Post, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &orders); err != nil {
		return c.Status(500).SendString(err.Error())

	}

	// return orders list in JSON format
	return c.JSON(orders)
}

func getPost(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func createPost(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func updatePost(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}

func deletePost(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
	return nil
}
