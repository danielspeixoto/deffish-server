package topic

import (
	boundary "deffish-server/src/boundary/topic"
	"deffish-server/src/domain/topic"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Repo boundary.IRepository
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.handle(upload))
	router.GET("/", handler.handle(all))
	router.GET("/:id", handler.handle(topicById))
}

func (handler Router) handle(callback func(Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(
			newControllerDefaults(Presenter{
				Writer: c.Writer,
		}, handler.Repo), c)
	}
}

func upload(ctrl Controller, ctx *gin.Context)       { ctrl.Upload(ctx) }
func all(ctrl Controller, ctx *gin.Context)       {
	ctrl.Random(ctx)
}
func topicById(ctrl Controller, ctx *gin.Context) { ctrl.TopicById(ctx) }

func newControllerDefaults(presenter Presenter, repo boundary.IRepository) Controller {
	uploadTopic := topic.Upload{Repo: repo, Presenter: presenter}
	random := topic.Random{Repo:repo, Presenter:presenter, Max:10}
	byId := topic.ById{Repo:repo, Presenter:presenter}

	return Controller{
		UploadUseCase: uploadTopic,
		RandomUseCase:random,
		GetById: byId,
	}
}