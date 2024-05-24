package inputs

type GetFacilityDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type FacilityInput struct {
	Name string `json:"name" binding:"required"`
}
