package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type AdminFormatter struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Token string      `json:"token,omitempty"`
	User  models.User `json:"user"`
}

func FormatAuthAdmin(admin models.Admin, token string) AdminFormatter {
	formatter := AdminFormatter{
		ID:    admin.ID,
		Name:  admin.Name,
		Token: token,
		User:  admin.User,
	}

	return formatter
}

type CustomerFormatter struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	Token       string      `json:"token,omitempty"`
	User        models.User `json:"user"`
}

func FormatAuthCustomer(customer models.Customer, token string) CustomerFormatter {
	formatter := CustomerFormatter{
		ID:          customer.ID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Token:       token,
		User:        customer.User,
	}

	return formatter
}

type InstructorFormatter struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	EmployeeCode string      `json:"employee_code"`
	Token        string      `json:"token,omitempty"`
	User         models.User `json:"user"`
}

func FormatAuthInstructor(instructor models.Instructor, token string) InstructorFormatter {
	formatter := InstructorFormatter{
		ID:           instructor.ID,
		Name:         instructor.Name,
		EmployeeCode: instructor.EmployeeCode,
		Token:        token,
		User:         instructor.User,
	}

	return formatter
}

type StaffFormatter struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	EmployeeCode string      `json:"employee_code"`
	Token        string      `json:"token,omitempty"`
	User         models.User `json:"user"`
}

func FormatAuthStaff(staff models.Staff, token string) StaffFormatter {
	formatter := StaffFormatter{
		ID:           staff.ID,
		Name:         staff.Name,
		EmployeeCode: staff.EmployeeCode,
		Token:        token,
		User:         staff.User,
	}

	return formatter
}
