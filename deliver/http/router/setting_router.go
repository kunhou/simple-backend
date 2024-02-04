package router

import (
	"net/http"

	"gorm.io/datatypes"

	"github/kunhou/simple-backend/deliver/http/schema"
	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/pkg/ginhelper"
	"github/kunhou/simple-backend/usecase/setting"
)

type SettingRouter struct {
	s *setting.SettingUsecase
}

func NewSettingRouter(s *setting.SettingUsecase) *SettingRouter {
	return &SettingRouter{
		s: s,
	}
}

func (router *SettingRouter) RegisterRouter(r *ginhelper.RouterGroup) {
	r.AddRouter(http.MethodPost, "/settings/", router.CreateSetting)
	r.AddRouter(http.MethodGet, "/settings/:setting_name", router.GetSettingByName)
	r.AddRouter(http.MethodPatch, "/settings/:setting_name", router.UpdateSetting)
	r.AddRouter(http.MethodDelete, "/settings/:setting_name", router.DeleteSetting)
}

func (router *SettingRouter) GetSettingByName(ctx *ginhelper.Context) {
	name := ctx.Param("setting_name")
	data, err := router.s.GetSettingByName(ctx, name)
	ctx.RespondWithError(err, data)
}

func (router *SettingRouter) CreateSetting(ctx *ginhelper.Context) {
	var setting schema.CreateSettingReq
	if ctx.BindAndCheck(&setting) {
		return
	}
	v, err := setting.Value.MarshalJSON()
	if err != nil {
		ctx.RespondWithError(err, nil)
		return
	}

	s := entity.Setting{
		Name:  &setting.Name,
		Value: &datatypes.JSON{},
	}
	if err := s.Value.Scan(v); err != nil {
		ctx.RespondWithError(err, nil)
		return
	}

	data, err := router.s.CreateSetting(ctx, &s)
	ctx.RespondWithError(err, data)
}

func (router *SettingRouter) UpdateSetting(ctx *ginhelper.Context) {
	name := ctx.Param("setting_name")

	var setting schema.UpdateSettingReq
	if ctx.BindAndCheck(&setting) {
		return
	}
	v, err := setting.Value.MarshalJSON()
	if err != nil {
		ctx.RespondWithError(err, nil)
		return
	}

	s := entity.Setting{
		Value: &datatypes.JSON{},
	}
	if err := s.Value.Scan(v); err != nil {
		ctx.RespondWithError(err, nil)
		return
	}

	data, err := router.s.UpdateSettingByName(ctx, name, &s)
	ctx.RespondWithError(err, data)
}

func (router *SettingRouter) DeleteSetting(ctx *ginhelper.Context) {
	name := ctx.Param("setting_name")
	err := router.s.DeleteSettingByName(ctx, name)

	ctx.RespondWithError(err, nil)
}
