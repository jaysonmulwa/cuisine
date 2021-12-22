package model

type MasterClassItem struct {
	ID                 string `json:"id,omitempty" bson:"_id,omitempty"`
	MasterClass_id     string `json:"master_class_id" bson:"master_class_id"`
	Item_order_no      int    `json:"item_order_no" bson:"item_order_no"`
	Item_title         string `json:"item_title" bson:"item_title"`
	Item_resource_url  string `json:"item_resource" bson:"item_resource"`
	Item_resource_type string `json:"item_resource_type " bson:"item_resource_type "`
	Item_content       string `json:"item_content" bson:"item_content"`
	Item_active        int    `json:"item_active" bson:"item_active"`
	Item_created_date  string `json:"item_created_date" bson:"item_created_date"`
	Item_updated_date  string `json:"item_updated_date" bson:"item_updated_date"`
}
