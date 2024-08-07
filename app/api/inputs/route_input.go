package inputs

type GetRouteDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type RouteInput struct {
	Name string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	// StartingRoute string `json:"starting_route" binding:"required"`
	// FinalRoute string `json:"final_route" binding:"required"`
	// Duration uint8 `json:"duration" binding:"required"`
	// Distance uint16 `json:"distance" binding:"required"`
}
