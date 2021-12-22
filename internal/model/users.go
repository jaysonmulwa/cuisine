package model

type User struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
	First_name   string `json:"first_name" bson:"first_name"`
	Last_name    string `json:"last_name" bson:"last_name"`
	Telephone    string `json:"telephone" bson:"telephone"`
	Country      string `json:"country" bson:"country"`
	City         string `json:"city" bson:"city"`
	Profile_pic  string `json:"profile_pic" bson:"profile_pic"`
	Verified_ind int    `json:"verified_ind" bson:"verified_ind"`
	Premium_ind  int    `json:"premium_ind" bson:"premium_ind"`
	Active_ind   int    `json:"active_ind" bson:"active_ind"`
	Created_date string `json:"created_date" bson:"created_date"`
	Updated_date string `json:"updated_date" bson:"updated_date"`
}
