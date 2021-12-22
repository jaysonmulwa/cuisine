package posts

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	"github.com/jaysonmulwa/cuisine/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := db.GetMongo().Db.Collection("posts").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var posts []model.Post = make([]model.Post, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())

	}

	// return orders list in JSON format
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {

	id := c.Params("id")

	result := model.Post{}
	postId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "_id", Value: postId}}
	err = db.GetMongo().Db.Collection("posts").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)

}

func GetPostsByUser(c *fiber.Ctx) error {

	id := c.Params("user_id")

	userId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "user_id", Value: userId}}
	cursor, err := db.GetMongo().Db.Collection("posts").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var posts []model.Post = make([]model.Post, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(posts)

}

func CreatePost(c *fiber.Ctx) error {

	post := new(model.Post)

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

	insertionResult, _ := db.GetMongo().Db.Collection("posts").InsertOne(c.Context(), newPost)

	fmt.Println(insertionResult)

	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	postId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	post := new(model.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
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

	err = db.GetMongo().Db.Collection("posts").FindOneAndUpdate(c.Context(), bson.M{"_id": postId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	post.ID = id
	return c.Status(200).JSON(post)

}

func DeletePost(c *fiber.Ctx) error {
	//Update the record -- ActiveInd = 0
	c.SendString("Hello, World!")
	return nil
}
