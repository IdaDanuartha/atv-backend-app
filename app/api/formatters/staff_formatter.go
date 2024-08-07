package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

func FormatStaff(staff models.Staff) models.Staff {
	staffFormatter := models.Staff{}
	staffFormatter.ID = staff.ID
	staffFormatter.Name = staff.Name
	staffFormatter.EmployeeCode = staff.EmployeeCode
	staffFormatter.CreatedAt = staff.CreatedAt
	staffFormatter.UpdatedAt = staff.UpdatedAt
	staffFormatter.DeletedAt = staff.DeletedAt

	staffFormatter.User.ID = staff.User.ID
	staffFormatter.User.Username = staff.User.Username
	staffFormatter.User.Email = staff.User.Email
	staffFormatter.User.Password = staff.User.Password
	staffFormatter.User.Role = staff.User.Role
	staffFormatter.User.CreatedAt = staff.User.CreatedAt
	staffFormatter.User.UpdatedAt = staff.User.UpdatedAt
	staffFormatter.User.DeletedAt = staff.User.DeletedAt

	return staffFormatter
}

func FormatStaffs(staffs []models.Staff) []models.Staff {
	staffsFormatter := []models.Staff{}

	for _, staff := range staffs {
		staff := FormatStaff(staff)
		staffsFormatter = append(staffsFormatter, staff)
	}

	return staffsFormatter
}
