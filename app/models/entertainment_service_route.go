package models

// Entertainment Service Route Model
type EntertainmentServiceRoute struct {
	Base
	EntertainmentServiceID string `gorm:"type:varchar(100);primaryKey;foreignKey:EntertainmentServiceID" json:"entertainment_service_id,omitempty"`
	Route                  Route  `gorm:"foreignKey:RouteID" json:"route"`
	RouteID                string `gorm:"type:varchar(100);primaryKey;foreignKey:RouteID" json:"route_id,omitempty"`
}

// TableName method sets table name for Entertainment Service Instructor model
func (entertainmentServiceRoute *EntertainmentServiceRoute) TableName() string {
	return "entertainment_service_routes"
}
