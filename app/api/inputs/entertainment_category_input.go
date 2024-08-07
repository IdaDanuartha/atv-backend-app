package inputs

type GetEntertainmentCategoryDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentCategoryInput struct {
	Name string `json:"name" binding:"required"`
}