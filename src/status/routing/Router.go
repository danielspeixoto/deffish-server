package routing

import (
	"deffish-server/src/status/domain"
	"deffish-server/src/status/presentation"
	"github.com/gin-gonic/gin"
)

type Router struct {}

func NewRouter() Router {
	return Router{}
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.handle(info))
}

func (handler Router) handle(callback func(presentation.Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(newControllerDefaults(presentation.Presenter{
			Writer: c.Writer,
		}), c)
	}
}

func info(ctrl presentation.Controller, ctx *gin.Context)       { ctrl.Status(ctx.Request) }

func newControllerDefaults(presenter presentation.Presenter) presentation.Controller {
	info := domain.Info{Presenter: presenter}

	return presentation.Controller{
		StatusUseCase:info,
	}
}