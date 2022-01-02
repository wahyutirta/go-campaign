package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
	// deta menggunakan type interface agar lebih flexibel dalam menyajikan data
}

type Metadata struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	metadata := Metadata{
		Message: message,
		Code:    code,
		Status:  status}

	jsonres := Response{
		Metadata: metadata,
		Data:     data,
	}

	return jsonres
}

func FormatValidationError(err error) []string {
	var Errors []string
	for _, e := range err.(validator.ValidationErrors) {
		Errors = append(Errors, e.Error())
	}
	return Errors
}
