package server

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"github/kunhou/simple-backend/deliver/http/router"
)

func NewHTTPServer(debug bool, settingRouter *router.SettingRouter) *gin.Engine {
	ginMode := gin.ReleaseMode
	if debug {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	r := gin.New()
	initVersionInfo()

	// register prometheus metrics
	p := ginprometheus.NewPrometheus("gin")
	p.MetricsPath = "/_metrics"
	p.Use(r)

	r.GET("/_health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	rootGroup := r.Group("")
	settingRouter.RegisterRouter(rootGroup)

	return r
}
