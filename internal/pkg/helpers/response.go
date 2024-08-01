package helpers

type ResponseData struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func Response(message string, code int, data any) ResponseData {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  "success",
	}

	if code >= 400 {
		meta.Status = "failed"
	}
	response := ResponseData{
		Meta: meta,
		Data: data,
	}
	return response
}
