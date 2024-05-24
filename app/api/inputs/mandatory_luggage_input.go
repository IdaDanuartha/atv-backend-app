package inputs

type GetMandatoryLuggageDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type MandatoryLuggageInput struct {
	Name string `json:"name" binding:"required"`
}
