package model

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
