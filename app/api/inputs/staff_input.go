package inputs

type GetStaffDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type StaffInput struct {
	Name         string `json:"name" binding:"required"`
	EmployeeCode string `json:"employee_code" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}
