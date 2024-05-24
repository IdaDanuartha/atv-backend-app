package inputs

type GetEntertainmentPackageDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentPackageInput struct {
	Name string `json:"name" binding:"required"`
}
