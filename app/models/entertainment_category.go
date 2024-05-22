package models

// Entertainment Catetgory Model
type EntertainmentCategory struct {
	Base
	Name         string    `gorm:"size:100" json:"name"`
}

// TableName method sets table name for Bus model
func (entertainmentCategory *EntertainmentCategory) TableName() string {
	return "entertainment_categories"
}

//ResponseMap -> response map method of Entertainment Category
func (EntertainmentCategory *EntertainmentCategory) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = EntertainmentCategory.ID
	resp["name"] = EntertainmentCategory.Name
	resp["created_at"] = EntertainmentCategory.CreatedAt
	resp["updated_at"] = EntertainmentCategory.UpdatedAt
	resp["deleted_at"] = EntertainmentCategory.DeletedAt

	return resp
}
