package http

import (
	"context"
	"fmt"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Recipe_name  string `json:"recipe_name" bson:"recipe_name"`
	Ingredients  string `json:"ingredients" bson:"ingredients"`
	Instructions string `json:"instructions" bson:"instructions"`
	Origin       string `json:"origin" bson:"origin"`
	Recipe_image string `json:"recipe_image" bson:"recipe_image"`
	Post_date    string `json:"post_date" bson:"post_date"`
	Post_type    string `json:"post_type" bson:"post_type"`
	User_id      int    `json:"user_id" bson:"user_id"`
	Active_Ind   int    `json:"active_ind" bson:"active_ind"`
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
	app.Get("/posts/:user_id", getPostsByUser)
	app.Post("/posts", createPost)
	app.Put("/posts/:id", updatePost)
	app.Delete("/posts/:id", deletePost)

	app.Listen(":4000")

}

func getAll(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("posts").Find(c.Context(), query)
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

	id := c.Params("id")

	result := Post{}
	postId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "_id", Value: postId}}
	err = mg.Db.Collection("posts").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)

}

func getPostsByUser(c *fiber.Ctx) error {

	id := c.Params("user_id")

	userId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "user_id", Value: userId}}
	cursor, err := mg.Db.Collection("posts").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var posts []Post = make([]Post, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(posts)

}

func createPost(c *fiber.Ctx) error {

	post := new(Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newPost := bson.D{
		{Key: "recipe_name", Value: post.Recipe_name},
		{Key: "ingredients", Value: post.Ingredients},
		{Key: "instructions", Value: post.Instructions},
		{Key: "origin", Value: post.Origin},
		{Key: "recipe_image", Value: post.Recipe_image},
		{Key: "post_date", Value: post.Post_date},
		{Key: "post_type", Value: post.Post_type},
		{Key: "user_id", Value: post.User_id},
		{Key: "active_ind", Value: 1},
	}

	insertionResult, _ := mg.Db.Collection("posts").InsertOne(c.Context(), newPost)

	fmt.Println(insertionResult)

	return c.JSON(post)
}

func updatePost(c *fiber.Ctx) error {

	id := c.Params("id")

	post := new(Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// the provided ID might be invalid ObjectID
	if err != nil {
	c.SendString("Hello, World!")
	return nil
}

func deletePost(c *fiber.Ctx) error {
	//Update the record -- ActiveInd = 0
	c.SendString("Hello, World!")
	return nil
}
