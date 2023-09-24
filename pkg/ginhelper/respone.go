package ginhelper

import (
	"github/kunhou/simple-backend/pkg/errcraft"
	"github/kunhou/simple-backend/pkg/reason"
)

type ResponseBody struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code   int           `json:"code"`
	Reason reason.Reason `json:"reason"`
}

func NewRespBodyData(code int, reason reason.Reason, data interface{}) *ResponseBody {
	return &ResponseBody{
		Meta: &Meta{
			Code:   code,
			Reason: reason,
		},
		Data: data,
	}
}

// NewRespBody new response body
func NewRespBody(code int, reason reason.Reason) *ResponseBody {
	return &ResponseBody{
		Meta: &Meta{
			Code:   code,
			Reason: reason,
		},
	}
}

// NewRespBodyFromError new response body from error
func NewRespBodyFromError(e *errcraft.Error) *ResponseBody {
	return &ResponseBody{
		Meta: &Meta{
			Code:   e.Code,
			Reason: e.Reason,
		},
	}
}
