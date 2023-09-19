package ginhelper

type ResponseBody struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code" snake:"code"`
	Message string `json:"message" snake:"message"`
}

func NewRespBodyData(code int, reason string, data interface{}) *ResponseBody {
	return &ResponseBody{
		Meta: &Meta{
			Code:    code,
			Message: reason,
		},
		Data: data,
	}
}
