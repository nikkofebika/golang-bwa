package helper

type apiResponse struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

type meta struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func ApiResponse(code uint16, message string, data any) apiResponse {
	meta := meta{Code: code, Message: message}

	return apiResponse{Meta: meta, Data: data}
}
