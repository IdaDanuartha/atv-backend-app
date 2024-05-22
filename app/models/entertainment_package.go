package models

// Entertainment Package Model
type EntertainmentPackage struct {
	Base
	Name string `gorm:"size:100" json:"name"`
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
	resp["created_at"] = EntertainmentPackage.CreatedAt
	resp["updated_at"] = EntertainmentPackage.UpdatedAt
	resp["deleted_at"] = EntertainmentPackage.DeletedAt

	return resp
}
