package models

import "time"

// Entertainment Package Model
type EntertainmentPackage struct {
	Base
	Name string `gorm:"size:100" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Price int32 `json:"price"`
	ExpiredAt time.Time `json:"expired_at"`
}

// TableName method sets table name for Bus model
func (entertainmentPackage *EntertainmentPackage) TableName() string {
	return "entertainment_packages"
}

// ResponseMap -> response map method of Entertainment Category
func (EntertainmentPackage *EntertainmentPackage) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = EntertainmentPackage.ID
	resp["name"] = EntertainmentPackage.Name
	resp["description"] = EntertainmentPackage.Description
	resp["price"] = EntertainmentPackage.Price
	resp["expired_at"] = EntertainmentPackage.ExpiredAt
	resp["created_at"] = EntertainmentPackage.CreatedAt
	resp["updated_at"] = EntertainmentPackage.UpdatedAt
	resp["deleted_at"] = EntertainmentPackage.DeletedAt

	return resp
}
