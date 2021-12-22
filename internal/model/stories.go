package model

type Story struct {
	ID                 string `json:"id,omitempty" bson:"_id,omitempty"`
	Story_name         string `json:"story_name" bson:"story_name"`
	Story_resource     string `json:"story_resource" bson:"story_resource"`
	Story_url          string `json:"story_url" bson:"story_url"`
	Story_type         string `json:"story_type" bson:"story_type"`
	Story_active       int    `json:"story_active" bson:"story_active"`
	Story_created_date string `json:"story_created_date" bson:"story_created_date"`
	Story_updated_date string `json:"story_updated_date" bson:"story_updated_date"`
	User_id            string `json:"user_id" bson:"user_id"`
}
