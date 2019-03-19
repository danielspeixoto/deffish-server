package essay

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Controller func(Presenter)Controller
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.handle(
		func(ctrl Controller, ctx *gin.Context) {
			ctrl.Upload(ctx)
		}))
	router.GET("/", handler.handle(
		func(ctrl Controller, ctx *gin.Context) {
			ctrl.Random(ctx)
		}))
	router.POST("/:id/comment", handler.handle(
		func (ctrl Controller, ctx *gin.Context) {
			ctrl.Comment(ctx)
		}))
	router.GET("/:id", handler.handle(
		func (ctrl Controller, ctx *gin.Context) {
			ctrl.EssayById(ctx)
		}))
}

func (handler Router) handle(callback func(Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(handler.Controller, c)
	}
}



