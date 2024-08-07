package inputs

type GetCustomerDetailInput struct {
	ID string `uri:"id" binding:"required"`
}