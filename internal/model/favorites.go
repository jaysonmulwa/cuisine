package model

type Favorite struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Post_id      int    `json:"post_id" bson:"post_id"`
	User_id      int    `json:"user_id" bson:"user_id"`
	Created_date string `json:"created_date" bson:"created_date"`
	Updated_date string `json:"updated_date" bson:"updated_date"`
	Active_ind   int    `json:"active_ind" bson:"active_ind"`
}
