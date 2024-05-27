package models

// Instructor Model
type Instructor struct {
	Base
	Name         string `gorm:"size:100" json:"name"`
	EmployeeCode string `gorm:"size:50" json:"employee_code"`
	UserID       string `gorm:"primaryKey;foreignKey:UserID"`
	User         User   `gorm:"foreignKey:UserID"`
}

// TableName method sets table name for Instructor model
func (Instructor *Instructor) TableName() string {
	return "instructors"
}

// ResponseMap -> response map method of Instructor
func (Instructor *Instructor) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Instructor.ID
	resp["name"] = Instructor.Name
	resp["created_at"] = Instructor.CreatedAt
	resp["updated_at"] = Instructor.UpdatedAt
	resp["deleted_at"] = Instructor.DeletedAt

	return resp
}
