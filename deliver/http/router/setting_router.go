package router

import (
	"github.com/gin-gonic/gin"
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

func (router *SettingRouter) RegisterRouter(r *gin.RouterGroup) {
	r.POST("/settings/", router.CreateSetting)
	r.GET("/settings/:setting_name", router.GetSettingByName)
	r.PATCH("/settings/:setting_name", router.UpdateSetting)
	r.DELETE("/settings/:setting_name", router.DeleteSetting)
}

func (router *SettingRouter) GetSettingByName(ctx *gin.Context) {
	name := ctx.Param("setting_name")
	data, err := router.s.GetSettingByName(ctx, name)
	ginhelper.RespondWithError(ctx, err, data)
}

func (router *SettingRouter) CreateSetting(ctx *gin.Context) {
	var setting schema.CreateSettingReq
	if ginhelper.BindAndCheck(ctx, &setting) {
		return
	}
	v, err := setting.Value.MarshalJSON()
	if err != nil {
		ginhelper.RespondWithError(ctx, err, nil)
		return
	}

	s := entity.Setting{
		Name:  &setting.Name,
		Value: &datatypes.JSON{},
	}
	if err := s.Value.Scan(v); err != nil {
		ginhelper.RespondWithError(ctx, err, nil)
		return
	}

	data, err := router.s.CreateSetting(ctx, &s)
	ginhelper.RespondWithError(ctx, err, data)
}

func (router *SettingRouter) UpdateSetting(ctx *gin.Context) {
	name := ctx.Param("setting_name")

	var setting schema.UpdateSettingReq
	if ginhelper.BindAndCheck(ctx, &setting) {
		return
	}
	v, err := setting.Value.MarshalJSON()
	if err != nil {
		ginhelper.RespondWithError(ctx, err, nil)
		return
	}

	s := entity.Setting{
		Value: &datatypes.JSON{},
	}
	if err := s.Value.Scan(v); err != nil {
		ginhelper.RespondWithError(ctx, err, nil)
		return
	}

	data, err := router.s.UpdateSettingByName(ctx, name, &s)
	ginhelper.RespondWithError(ctx, err, data)
}

func (router *SettingRouter) DeleteSetting(ctx *gin.Context) {
	name := ctx.Param("setting_name")
	err := router.s.DeleteSettingByName(ctx, name)

	ginhelper.RespondWithError(ctx, err, nil)
}
