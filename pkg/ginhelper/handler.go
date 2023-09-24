package ginhelper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/kunhou/simple-backend/pkg/errcraft"
	"github/kunhou/simple-backend/pkg/reason"
)

// BindAndCheck attempts to bind the request data and checks for errors.
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	if err := ctx.ShouldBind(data); err != nil {
		logError(fmt.Sprintf("Failed to bind data: %s", err.Error()))
		RespondWithError(ctx, errcraft.New(http.StatusBadRequest, reason.InvalidRequest), nil)
		return true
	}
	return false
}

// RespondWithError checks the error and responds appropriately.
func RespondWithError(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		respondWithSuccess(ctx, data)
		return
	}

	originErr := errcraft.Unwrap(err)
	if customErr, ok := originErr.(*errcraft.Error); ok {
		handleCustomError(ctx, customErr, data)
		return
	}

	logErrorAndRespondWithUnknown(ctx, err)
}

func respondWithSuccess(ctx *gin.Context, data interface{}) {
	response := NewRespBodyData(http.StatusOK, reason.Success, data)
	ctx.JSON(http.StatusOK, response)
}

func handleCustomError(ctx *gin.Context, err *errcraft.Error, data interface{}) {
	if errcraft.IsInternalServer(err) {
		logError(fmt.Sprintf("Internal server error: %s", err.Error()))
	}

	respBody := NewRespBodyFromError(err)
	if data != nil {
		respBody.Data = data
	}
	ctx.JSON(err.Code, respBody)
}

func logErrorAndRespondWithUnknown(ctx *gin.Context, err error) {
	logError(fmt.Sprintf("Unknown error: %s", err.Error()))
	response := NewRespBody(http.StatusInternalServerError, reason.Unknown)
	ctx.JSON(http.StatusInternalServerError, response)
}

func logError(message string) {
	log.Printf("%s\n%s", message, errcraft.CaptureStack(2, 5))
}
