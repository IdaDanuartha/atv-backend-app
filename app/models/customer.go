package models

// Customer Model
type Customer struct {
	Base
	Name 	string `gorm:"size:100" json:"name"`
	PhoneNumber 	string `gorm:"size:100" json:"phone_number"`
	UserID  string    `gorm:"type:varchar(100);primaryKey;foreignKey:UserID" json:"user_id,omitempty"`
  	User    User   `gorm:"foreignKey:UserID" json:"user"`
}

// TableName method sets table name for Customer model
func (Customer *Customer) TableName() string {
	return "customers"
}