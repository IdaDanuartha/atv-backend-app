package models

// Entertainment Package Model
type Blog struct {
	Base
	Title       string  `gorm:"size:100" json:"title"`
	Slug        string  `gorm:"size:100" json:"slug"`
	Description string  `gorm:"type:text" json:"description"`
	Content     string  `gorm:"type:text" json:"content"`
	ImagePath   *string `gorm:"size:150;" json:"image_path"`
}

// TableName method sets table name for Bus model
func (blog *Blog) TableName() string {
	return "blogs"
}
