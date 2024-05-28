package inputs

type GetEntertainmentPackageDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type EntertainmentPackageInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       string `json:"price" binding:"required"`
	ExpiredAt   string `json:"expired_at" binding:"required"`
}
