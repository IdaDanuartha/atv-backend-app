package inputs

type GetEntertainmentServiceDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentServiceInput struct {
	Name                    string `json:"name" binding:"required"`
	Price                   int32  `json:"price" binding:"required"`
	EntertainmentCategoryID string `json:"entertainment_category_id" binding:"required"`
	RouteID                 string `json:"route_id" binding:"required"`
}
