package ginhelper

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, "success", data))
		return
	}

	ctx.JSON(http.StatusInternalServerError, &ResponseBody{
		Meta: &Meta{
			Code:    500,
			Message: err.Error(),
		},
	})
}

func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	if err := ctx.ShouldBind(data); err != nil {
		log.Printf("http_handle BindAndCheck fail, %s", err.Error())
		Response(ctx, err, nil)
		return true
	}

	return false
}
