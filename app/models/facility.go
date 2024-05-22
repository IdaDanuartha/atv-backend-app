package models

// Facility Model
type Facility struct {
	Base
	Name string `gorm:"size:100" json:"name"`
}

// TableName method sets table name for Bus model
func (Facility *Facility) TableName() string {
	return "facilities"
}

// ResponseMap -> response map method of Entertainment Category
func (Facility *Facility) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Facility.ID
	resp["name"] = Facility.Name
	resp["created_at"] = Facility.CreatedAt
	resp["updated_at"] = Facility.UpdatedAt
	resp["deleted_at"] = Facility.DeletedAt

	return resp
}
