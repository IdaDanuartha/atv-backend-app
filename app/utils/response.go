package utils

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message     string `json:"message"`
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Total       int64  `json:"total,omitempty"`
	CurrentPage int    `json:"current_page,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func APIResponseWithPagination(message string, code int, status string, total int64, currentPage int, pageSize int, data interface{}) Response {
	meta := Meta{
		Message:     message,
		Code:        code,
		Status:      status,
		Total:       total,
		CurrentPage: currentPage,
		PageSize:    10,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
