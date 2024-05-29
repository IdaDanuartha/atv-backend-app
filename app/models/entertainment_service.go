package models

// Entertainment Service Model
type EntertainmentService struct {
	Base
	Name        string    `gorm:"size:100" json:"name"`
	Price       int32     `json:"price"`

	EntertainmentCategoryID string `gorm:"type:varchar(100);primaryKey;foreignKey:EntertainmentCategoryID" json:"entertainment_category_id,omitempty"`
	EntertainmentCategory EntertainmentCategory `gorm:"foreignKey:EntertainmentCategoryID" json:"entertainment_category"`
	
	RouteID string `gorm:"type:varchar(100);primaryKey;foreignKey:RouteID" json:"route_id,omitempty"`
	Route Route `gorm:"foreignKey:RouteID" json:"route"`
}

// TableName method sets table name for Bus model
func (entertainmentService *EntertainmentService) TableName() string {
	return "entertainment_services"
}