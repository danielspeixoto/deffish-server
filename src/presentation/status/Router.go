package status

import (
	"deffish-server/src/domain/status"
	"github.com/gin-gonic/gin"
)

type Router struct {}

func (handler Router) Route(router *gin.RouterGroup) {
	router.GET("/", handler.handle(info))
}

func (handler Router) handle(callback func(Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(newControllerDefaults(Presenter{
			Writer: c.Writer,
		}), c)
	}
}

func info(ctrl Controller, ctx *gin.Context)       {
	ctrl.Status(ctx.Request)
}

func newControllerDefaults(presenter Presenter) Controller {
	info := status.Info{Presenter: presenter}

	return Controller{
		StatusUseCase:info,
	}
}