package model

type Stream struct {
	ID                  string `json:"id,omitempty" bson:"_id,omitempty"`
	Stream_name         string `json:"stream_name" bson:"stream_name"`
	Stream_resource     string `json:"stream_resource" bson:"stream_resource"`
	Stream_url          string `json:"stream_url" bson:"stream_url"`
	Stream_type         string `json:"stream_type" bson:"stream_type"`
	User_id             string `json:"user_id" bson:"user_id"`
	Stream_start_date   string `json:"stream_start_date" bson:"stream_start_date"`
	Stream_end_date     string `json:"stream_end_date" bson:"stream_end_date"`
	Stream_active_ind   int    `json:"stream_active_ind" bson:"stream_active_ind"`
	Stream_created_date string `json:"stream_created_date" bson:"stream_created_date"`
	Stream_updated_date string `json:"stream_updated_date" bson:"stream_updated_date"`
}
