package models

// Staff Model
type Staff struct {
	Base
	Name         string `gorm:"size:100" json:"name"`
	EmployeeCode string `gorm:"size:50" json:"employee_code"`
	UserID       string `gorm:"type:varchar(100);" json:"user_id,omitempty"`
	User         User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// TableName method sets table name for Staff model
func (Staff *Staff) TableName() string {
	return "staff"
}
