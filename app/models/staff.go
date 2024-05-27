package models

// Staff Model
type Staff struct {
	Base
	Name         string `gorm:"size:100" json:"name"`
	EmployeeCode string `gorm:"size:50" json:"employee_code"`
	UserID       string `gorm:"type:varchar(100);primaryKey;foreignKey:UserID"`
	User         User   `gorm:"foreignKey:UserID"`
}

// TableName method sets table name for Staff model
func (Staff *Staff) TableName() string {
	return "staff"
}

// ResponseMap -> response map method of Staff
func (Staff *Staff) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Staff.ID
	resp["name"] = Staff.Name
	resp["created_at"] = Staff.CreatedAt
	resp["updated_at"] = Staff.UpdatedAt
	resp["deleted_at"] = Staff.DeletedAt

	return resp
}
