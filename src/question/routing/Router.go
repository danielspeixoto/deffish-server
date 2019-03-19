package routing

import (
	"deffish-server/src/question"
	"deffish-server/src/question/domain"
	"deffish-server/src/question/presentation"
	"github.com/gin-gonic/gin"
)

type Router struct {
	repo question.IRepository
}

func NewRouter(repo question.IRepository) Router {
	return Router{
		repo,
	}
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.handle(upload))
	router.GET("/", handler.handle(random))
	router.GET("/:id", handler.handle(questionById))
}

func (handler Router) handle(callback func(presentation.Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(
			newControllerDefaults(presentation.Presenter{
				Writer: c.Writer,
		}, handler.repo), c)
	}
}

func upload(ctrl presentation.Controller, ctx *gin.Context)       { ctrl.Upload(ctx.Request) }
func random(ctrl presentation.Controller, ctx *gin.Context)       {
	ctrl.Random(ctx.Request)
}
func questionById(ctrl presentation.Controller, ctx *gin.Context) { ctrl.QuestionById(ctx) }

func newControllerDefaults(presenter presentation.Presenter, repo question.IRepository) presentation.Controller {
	uploadQuestion := domain.Upload{Repo: repo, Presenter: presenter}
	random := domain.Random{Repo:repo, Presenter:presenter, MaxQuestions:10}
	byId := domain.ById{Repo:repo, Presenter:presenter}

	return presentation.Controller{
		UploadUseCase: uploadQuestion,
		RandomUseCase:random,
		GetById: byId,
	}
}