package model

type MasterClass struct {
	ID                           string `json:"id,omitempty" bson:"_id,omitempty"`
	Masterclass_name             string `json:"masterclass_name" bson:"masterclass_name"`
	Masterclass_created_date     string `json:"masterclass_created_date" bson:"masterclass_created_date"`
	Masterclass_updated_date     string `json:"masterclass_updated_date" bson:"masterclass_updated_date"`
	Masterclass_active_ind       int    `json:"masterclass_active_ind" bson:"masterclass_active_ind"`
	Masterclass_start_date       string `json:"masterclass_start_date" bson:"masterclass_start_date"`
	Masterclass_end_date         string `json:"masterclass_end_date" bson:"masterclass_end_date"`
	Masterclass_description      string `json:"masterclass_description" bson:"masterclass_description"`
	Masterclass_image            string `json:"masterclass_image" bson:"masterclass_image"`
	Masterclass_url              string `json:"masterclass_url" bson:"masterclass_url"`
	Masterclass_type             string `json:"masterclass_type" bson:"masterclass_type"`
	Masterclass_verified_ind     int    `json:"masterclass_verified_ind" bson:"masterclass_verified_ind"`
	Masterclass_created_by       string `json:"masterclass_created_by" bson:"masterclass_created_by"`
	Masterclass_unit_price       string `json:"masterclass_unit_price" bson:"masterclass_unit_price"`
	Masterclass_currency         string `json:"masterclass_currency" bson:"masterclass_currency"`
	Masterclass_discount_percent string `json:"masterclass_discount_percent" bson:"masterclass_discount_percent"`
}
