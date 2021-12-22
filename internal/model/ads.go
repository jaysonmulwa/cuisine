package model

type Ad struct {
	ID              string `json:"id,omitempty" bson:"_id,omitempty"`
	Ad_type         string `json:"ad_type" bson:"ad_type"`
	Ad_text         string `json:"ad_text" bson:"ad_text"`
	Ad_image        string `json:"ad_image" bson:"ad_image"`
	Ad_url          string `json:"ad_url" bson:"ad_url"`
	Ad_start_date   string `json:"ad_start_date" bson:"ad_start_date"`
	Ad_end_date     string `json:"ad_end_date" bson:"ad_end_date"`
	Ad_active_ind   int    `json:"ad_active_ind" bson:"ad_active_ind"`
	Ad_created_date string `json:"ad_created_date" bson:"ad_created_date"`
	Ad_updated_date string `json:"ad_updated_date" bson:"ad_updated_date"`
}
