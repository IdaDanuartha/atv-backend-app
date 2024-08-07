package models

// Route Model
type Route struct {
	Base
	Name    string `gorm:"size:50" json:"name"`
	Address string `gorm:"size:100" json:"address"`
	// StartingRoute string `gorm:"size:50" json:"starting_route"`
	// FinalRoute string `gorm:"size:50" json:"final_route"`
	// Duration uint8 `json:"duration"`
	// Distance uint16 `json:"distance"`
}

// TableName method sets table name for Route model
func (route *Route) TableName() string {
	return "routes"
}
