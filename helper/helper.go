package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// Function for handle format response API
func ApiResponse(message string, code int, status string, data interface{}) Response {
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

// Function for handle format error
func FormatValidationError(err error) []string {
	// Create variable err with data type slice string
	var errors []string

	// Loop all error
	for _, e := range err.(validator.ValidationErrors) {
		// Append error message to variable errors
		errors = append(errors, e.Error())
	}

	return errors
}
