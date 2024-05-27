package models

// Instructor Model
type Instructor struct {
	Base
	Name         string `gorm:"size:100" json:"name"`
	EmployeeCode string `gorm:"size:50" json:"employee_code"`
	UserID       string `gorm:"type:varchar(100);primaryKey;foreignKey:UserID" json:"user_id"`
	User         User   `gorm:"foreignKey:UserID" json:"user"`
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
	resp["employee_code"] = Instructor.EmployeeCode
	resp["created_at"] = Instructor.CreatedAt
	resp["updated_at"] = Instructor.UpdatedAt
	resp["deleted_at"] = Instructor.DeletedAt

	return resp
}
