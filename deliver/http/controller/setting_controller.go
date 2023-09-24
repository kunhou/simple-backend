package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"

	"github/kunhou/simple-backend/deliver/http/schema"
	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/pkg/ginhelper"
	"github/kunhou/simple-backend/usecase/setting"
)

type SettingController struct {
	s *setting.SettingUsecase
}

func NewSettingController(s *setting.SettingUsecase) *SettingController {
	return &SettingController{
		s: s,
	}
}

func (d *SettingController) GetSettingByName(ctx *gin.Context) {
	name := ctx.Param("setting_name")
	data, err := d.s.GetSettingByName(ctx, name)
	ginhelper.RespondWithError(ctx, err, data)
}

func (d *SettingController) CreateSetting(ctx *gin.Context) {
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

	data, err := d.s.CreateSetting(ctx, &s)
	ginhelper.RespondWithError(ctx, err, data)
}

func (d *SettingController) UpdateSetting(ctx *gin.Context) {
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

	data, err := d.s.UpdateSettingByName(ctx, name, &s)
	ginhelper.RespondWithError(ctx, err, data)
}

func (d *SettingController) DeleteSetting(ctx *gin.Context) {
	name := ctx.Param("setting_name")
	err := d.s.DeleteSettingByName(ctx, name)

	ginhelper.RespondWithError(ctx, err, nil)
}
