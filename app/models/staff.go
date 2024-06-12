package models

// Staff Model
type Staff struct {
	Base
	Name         string `gorm:"size:100;uniqueIndex" json:"name"`
	EmployeeCode string `gorm:"size:50;uniqueIndex" json:"employee_code"`
	UserID       string `gorm:"type:varchar(100);primaryKey;foreignKey:UserID" json:"user_id,omitempty"`
	User         User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// TableName method sets table name for Staff model
func (Staff *Staff) TableName() string {
	return "staff"
}
