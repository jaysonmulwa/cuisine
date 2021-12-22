package model

type User struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
	First_name   string `json:"first_name" bson:"first_name"`
	Last_name    string `json:"last_name" bson:"last_name"`
	Active_ind   int    `json:"active_ind" bson:"active_ind"`
	Created_date string `json:"created_date" bson:"created_date"`
	Updated_date string `json:"updated_date" bson:"updated_date"`
}
