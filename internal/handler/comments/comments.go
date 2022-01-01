package comments

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	"github.com/jaysonmulwa/cuisine/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetComment(c *fiber.Ctx) error {

	id := c.Params("id")

	result := model.Post{}
	commentId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "_id", Value: commentId}}
	err = db.GetMongo().Db.Collection("comments").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)

}

func GetCommentsByPost(c *fiber.Ctx) error {

	id := c.Params("post_id")

	result := model.Post{}
	postId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "post_id", Value: postId}}
	err = db.GetMongo().Db.Collection("comments").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)

}

func GetCommentsByPostForUser(c *fiber.Ctx) error {

	id := c.Params("post_id")
	usr_id := c.Params("user_id")

	result := model.Post{}
	postId, err := primitive.ObjectIDFromHex(id)
	user_id, _ := primitive.ObjectIDFromHex(usr_id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "post_id", Value: postId}, {Key: "user_id", Value: user_id}}
	err = db.GetMongo().Db.Collection("comments").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON
	return c.JSON(result)

}

func CreateComment(c *fiber.Ctx) error {

	comment := new(model.Comment)

	if err := c.BodyParser(comment); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newPost := bson.D{
		{Key: "post_id", Value: comment.Post_id},
		{Key: "user_id", Value: comment.User_id},
		{Key: "parent_id", Value: comment.Parent_id},
		{Key: "comment_text", Value: comment.Comment_text},
		{Key: "created_date", Value: comment.Created_date},
		{Key: "updated_date", Value: comment.Updated_date},
		{Key: "active_ind", Value: 1},
		{Key: "flag_ind", Value: 0},
		{Key: "flagger_id", Value: 0},
	}

	insertionResult, _ := db.GetMongo().Db.Collection("comments").InsertOne(c.Context(), newPost)

	fmt.Println(insertionResult)

	return c.JSON(comment)

}

func UpdateComment(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	commentId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	comment := new(model.Comment)

	if err := c.BodyParser(comment); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "post_id", Value: comment.Post_id},
		{Key: "user_id", Value: comment.User_id},
		{Key: "parent_id", Value: comment.Parent_id},
		{Key: "comment_text", Value: comment.Comment_text},
		{Key: "created_date", Value: comment.Created_date},
		{Key: "updated_date", Value: comment.Updated_date},
		{Key: "active_ind", Value: comment.Active_ind},
		{Key: "flag_ind", Value: comment.Flag_ind},
		{Key: "flagger_id", Value: comment.Flagger_id},
	}

	err = db.GetMongo().Db.Collection("comments").FindOneAndUpdate(c.Context(), bson.M{"_id": commentId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	comment.ID = id
	return c.Status(200).JSON(comment)

}

func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")

	var err error

	commentId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	comment := new(model.Comment)

	if err := c.BodyParser(comment); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "active_ind", Value: 0},
	}

	err = db.GetMongo().Db.Collection("comments").FindOneAndUpdate(c.Context(), bson.M{"_id": commentId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	comment.ID = id
	return c.Status(200).JSON(comment)

}

func FlagComment(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	commentId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	comment := new(model.Comment)

	if err := c.BodyParser(comment); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "flag_ind", Value: comment.Flag_ind},
		{Key: "flagger_id", Value: comment.Flagger_id},
		{Key: "flag_date", Value: comment.Flag_date},
		{Key: "flag_reason", Value: comment.Flag_reason},
	}

	err = db.GetMongo().Db.Collection("comments").FindOneAndUpdate(c.Context(), bson.M{"_id": commentId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	comment.ID = id
	return c.Status(200).JSON(comment)

}
