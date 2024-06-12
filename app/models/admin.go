package models

// Admin Model
type Admin struct {
	Base
	Name   string `gorm:"size:100;uniqueIndex" json:"name"`
	UserID string `gorm:"type:varchar(100);primaryKey;foreignKey:UserID" json:"user_id,omitempty"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// TableName method sets table name for Admin model
func (Admin *Admin) TableName() string {
	return "admins"
}
