package models

// Route Model
type Route struct {
	Base
	StartingRoute string `gorm:"size:50" json:"starting_route"`
	FinalRoute string `gorm:"size:50" json:"final_route"`
	Duration uint8 `json:"duration"`
	Distance uint16 `json:"distance"`
}

// TableName method sets table name for Route model
func (route *Route) TableName() string {
	return "routes"
}

// ResponseMap -> response map method of Route
func (route *Route) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = route.ID
	resp["starting_route"] = route.StartingRoute
	resp["final_route"] = route.FinalRoute
	resp["duration"] = route.Duration
	resp["distance"] = route.Distance
	resp["created_at"] = route.CreatedAt
	resp["updated_at"] = route.UpdatedAt
	resp["deleted_at"] = route.DeletedAt

	return resp
}
