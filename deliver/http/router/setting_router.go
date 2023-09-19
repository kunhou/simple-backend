package router

import (
	"github.com/gin-gonic/gin"

	"github/kunhou/simple-backend/deliver/http/controller"
)

type SettingRouter struct {
	settingController *controller.SettingController
}

func NewSettingRouter(settingController *controller.SettingController) *SettingRouter {
	return &SettingRouter{
		settingController: settingController,
	}
}

func (s *SettingRouter) RegisterRouter(r *gin.RouterGroup) {
	r.POST("/settings", s.settingController.CreateSetting)
	r.GET("/settings/:setting_name", s.settingController.GetSettingByName)
	r.PATCH("/settings/:setting_name", s.settingController.UpdateSetting)
	r.DELETE("/settings/:setting_name", s.settingController.DeleteSetting)
}
