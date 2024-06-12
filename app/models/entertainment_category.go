package models

// Entertainment Category Model
type EntertainmentCategory struct {
	Base
	Name string `gorm:"size:100;uniqueIndex" json:"name"`
}

// TableName method sets table name for Bus model
func (entertainmentCategory *EntertainmentCategory) TableName() string {
	return "entertainment_categories"
}
