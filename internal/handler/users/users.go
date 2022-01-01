package users

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	"github.com/jaysonmulwa/cuisine/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(c *fiber.Ctx) error {

	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := db.GetMongo().Db.Collection("users").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []model.User = make([]model.User, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())

	}

	// return orders list in JSON format
	return c.JSON(users)

}

func GetUser(c *fiber.Ctx) error {

	id := c.Params("id")

	result := model.User{}
	userId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "_id", Value: userId}}
	err = db.GetMongo().Db.Collection("users").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)

}

func CreateUser(c *fiber.Ctx) error {

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newPost := bson.D{
		{Key: "username", Value: user.Username},
		{Key: "password", Value: user.Password},
		{Key: "email", Value: user.Email},
		{Key: "first_name", Value: user.First_name},
		{Key: "last_name", Value: user.Last_name},
		{Key: "telephone", Value: user.Telephone},
		{Key: "country", Value: user.Country},
		{Key: "city", Value: user.City},
		{Key: "profile_pic", Value: user.Profile_pic},
		{Key: "verified_ind", Value: user.Verified_ind},
		{Key: "premium_ind", Value: user.Premium_ind},
		{Key: "active_ind", Value: user.Active_ind},
		{Key: "created_date", Value: user.Created_date},
		{Key: "updated_date", Value: user.Updated_date},
	}

	insertionResult, _ := db.GetMongo().Db.Collection("users").InsertOne(c.Context(), newPost)

	fmt.Println(insertionResult)

	return c.JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "username", Value: user.Username},
		{Key: "password", Value: user.Password},
		{Key: "email", Value: user.Email},
		{Key: "first_name", Value: user.First_name},
		{Key: "last_name", Value: user.Last_name},
		{Key: "telephone", Value: user.Telephone},
		{Key: "country", Value: user.Country},
		{Key: "city", Value: user.City},
		{Key: "profile_pic", Value: user.Profile_pic},
		{Key: "verified_ind", Value: user.Verified_ind},
		{Key: "premium_ind", Value: user.Premium_ind},
		{Key: "active_ind", Value: user.Active_ind},
		{Key: "created_date", Value: user.Created_date},
		{Key: "updated_date", Value: user.Updated_date},
	}

	err = db.GetMongo().Db.Collection("users").FindOneAndUpdate(c.Context(), bson.M{"_id": userId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	user.ID = id
	return c.Status(200).JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "active_ind", Value: 0},
	}

	err = db.GetMongo().Db.Collection("users").FindOneAndUpdate(c.Context(), bson.M{"_id": userId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	user.ID = id
	return c.Status(200).JSON(user)

}
