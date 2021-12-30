package helper

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
