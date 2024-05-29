package models

// Instructor Model
type Instructor struct {
	Base
	Name         string `gorm:"size:100" json:"name"`
	EmployeeCode string `gorm:"size:50" json:"employee_code"`
	UserID       string `gorm:"type:varchar(100);primaryKey;foreignKey:UserID" json:"user_id,omitempty"`
	User         User   `gorm:"foreignKey:UserID" json:"user"`
}

// TableName method sets table name for Instructor model
func (Instructor *Instructor) TableName() string {
	return "instructors"
}