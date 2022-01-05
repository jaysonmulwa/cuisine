package favorites

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	db "github.com/jaysonmulwa/cuisine/database"
	"github.com/jaysonmulwa/cuisine/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Todo! Get actual posts
func GetFavoritesByUser(c *fiber.Ctx) error {

	//id := c.Params("post_id")
	//usr := c.Params("user_id")

	//postId, err := primitive.ObjectIDFromHex(id)
	//userId, err := primitive.ObjectIDFromHex(usr)

	// the provided ID might be invalid ObjectID
	// if err != nil {
	// 	return c.SendStatus(400)
	// }

	// get all record as a cursor
	//query := bson.D{{Key: "post_id", Value: postId}, {Key: "user_id", Value: userId}}
	query := bson.D{{Key: "user_id", Value: 1}}
	cursor, err := db.GetMongo().Db.Collection("favorites").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var favorites []model.Favorite = make([]model.Favorite, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &favorites); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(favorites)
}

func DeleteFavorite(c *fiber.Ctx) error {

	id := c.Params("id")

	var err error

	favoriteId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	favorite := new(model.Favorite)

	if err := c.BodyParser(favorite); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	update := bson.D{
		{Key: "active_ind", Value: 0},
	}

	err = db.GetMongo().Db.Collection("favorites").FindOneAndUpdate(c.Context(), bson.M{"_id": favoriteId}, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}

	// return the updated order
	favorite.ID = id
	return c.Status(200).JSON(favorite)
}

func CreateFavorite(c *fiber.Ctx) error {

	like := new(model.Like)

	if err := c.BodyParser(like); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newPost := bson.D{
		{Key: "post_id", Value: like.Post_id},
		{Key: "user_id", Value: like.User_id},
		{Key: "created_date", Value: like.Created_date},
		{Key: "updated_date", Value: like.Updated_date},
		{Key: "active_ind", Value: 1},
	}

	insertionResult, _ := db.GetMongo().Db.Collection("likes").InsertOne(c.Context(), newPost)

	fmt.Println(insertionResult)

	return c.JSON(like)

}
