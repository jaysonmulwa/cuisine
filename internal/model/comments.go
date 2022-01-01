package model

type Comment struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	Post_id      string `json:"post_id" bson:"post_id"`
	User_id      string `json:"user_id" bson:"user_id"`
	Parent_id    string `json:"parent_id" bson:"parent_id"`
	Comment_text string `json:"comment_text" bson:"comment_text"`
	Created_date string `json:"created_date" bson:"created_date"`
	Updated_date string `json:"updated_date" bson:"updated_date"`
	Active_ind   int    `json:"active_ind" bson:"active_ind"`
	Flag_ind     int    `json:"flag_ind" bson:"flag_ind"`
	Flagger_id   string `json:"flagger_id" bson:"flagger_id"`
	Flag_reason  string `json:"flag_reason" bson:"flag_reason"`
	Flag_date    string `json:"flag_date" bson:"flag_date"`
}
