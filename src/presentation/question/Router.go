package question

import (
	"deffish-server/src/domain/question"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Repo question.IRepository
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.handle(upload))
	router.GET("/", handler.handle(random))
	router.GET("/:id", handler.handle(questionById))
}

func (handler Router) handle(callback func(Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(
			newControllerDefaults(Presenter{
				Writer: c.Writer,
		}, handler.Repo), c)
	}
}

func upload(ctrl Controller, ctx *gin.Context)       {
	ctrl.Upload(ctx.Request)
}
func random(ctrl Controller, ctx *gin.Context)       {
	ctrl.Random(ctx.Request)
}

func questionById(ctrl Controller, ctx *gin.Context) { ctrl.QuestionById(ctx) }

func newControllerDefaults(presenter Presenter, repo question.IRepository) Controller {
	uploadQuestion := question.Upload{Repo: repo, Presenter: presenter}
	random := question.Random{Repo:repo, Presenter:presenter, MaxQuestions:10}
	byId := question.ById{Repo:repo, Presenter:presenter}

	return Controller{
		UploadUseCase: uploadQuestion,
		RandomUseCase:random,
		GetById: byId,
	}
}