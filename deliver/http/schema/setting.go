package schema

import "encoding/json"

type CreateSettingReq struct {
	Name  string          `json:"name" binding:"required,gt=0,lte=35"`
	Value json.RawMessage `json:"value"`
}

type UpdateSettingReq struct {
	Value json.RawMessage `json:"value"`
}
