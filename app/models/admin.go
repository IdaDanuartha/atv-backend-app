package models

// Admin Model
type Admin struct {
	Base
	Name 	string `gorm:"size:100" json:"name"`
	UserID  string    `gorm:"primaryKey;foreignKey:UserID"`
  	User    User   `gorm:"foreignKey:UserID"`
}

// TableName method sets table name for Admin model
func (Admin *Admin) TableName() string {
	return "admins"
}

// ResponseMap -> response map method of Admin
func (Admin *Admin) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Admin.ID
	resp["name"] = Admin.Name
	resp["created_at"] = Admin.CreatedAt
	resp["updated_at"] = Admin.UpdatedAt
	resp["deleted_at"] = Admin.DeletedAt

	return resp
}
