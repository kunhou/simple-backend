package ginhelper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/kunhou/simple-backend/pkg/errcraft"
	"github/kunhou/simple-backend/pkg/reason"
)

type Context struct {
	*gin.Context
}

type RouterGroup struct {
	*gin.RouterGroup
}

func newContext(ctx *gin.Context) *Context {
	return &Context{ctx}
}

func NewRouterGroup(r *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{r}
}

type HandlerFunc func(*Context)

func (ctx *RouterGroup) AddRouter(httpMethod, relativePath string, handler HandlerFunc) {
	ctx.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
		handler(newContext(ctx))
	})
}

// BindAndCheck attempts to bind the request data and checks for errors.
func (ctx *Context) BindAndCheck(data interface{}) bool {
	if err := ctx.ShouldBind(data); err != nil {
		logError(fmt.Sprintf("Failed to bind data: %s", err.Error()))
		ctx.RespondWithError(errcraft.New(http.StatusBadRequest, reason.InvalidRequest), nil)
		return true
	}
	return false
}

// RespondWithError checks the error and responds appropriately.
func (ctx *Context) RespondWithError(err error, data interface{}) {
	if err == nil {
		ctx.respondWithSuccess(data)
		return
	}

	originErr := errcraft.Unwrap(err)
	if customErr, ok := originErr.(*errcraft.Error); ok {
		ctx.handleCustomError(customErr, data)
		return
	}

	ctx.logErrorAndRespondWithUnknown(err)
}

func (ctx *Context) respondWithSuccess(data interface{}) {
	response := NewRespBodyData(http.StatusOK, reason.Success, data)
	ctx.JSON(http.StatusOK, response)
}

func (ctx *Context) handleCustomError(err *errcraft.Error, data interface{}) {
	if errcraft.IsInternalServer(err) {
		logError(fmt.Sprintf("Internal server error: %s", err.Error()))
	}

	respBody := NewRespBodyFromError(err)
	if data != nil {
		respBody.Data = data
	}
	ctx.JSON(err.Code, respBody)
}

func (ctx *Context) logErrorAndRespondWithUnknown(err error) {
	logError(fmt.Sprintf("Unknown error: %s", err.Error()))
	response := NewRespBody(http.StatusInternalServerError, reason.Unknown)
	ctx.JSON(http.StatusInternalServerError, response)
}

func logError(message string) {
	log.Printf("%s\n%s", message, errcraft.CaptureStack(2, 5))
}
