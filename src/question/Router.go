package question

import (
	"deffish-server/src/question/data"
	"deffish-server/src/question/domain"
	"deffish-server/src/question/presentation"
	"deffish-server/src/status"
	"github.com/gin-gonic/gin"
)

type Router struct {
	repo domain.IRepository
}

func NewRouter(repo domain.IRepository) Router {
	return Router{
		repo,
	}
}

func NewRouterDefaults(uri string,
	database string,
	questionsCollection string) Router {
	return NewRouter(data.NewMongoRepository(
		uri, database, questionsCollection,
		))
}

func controller(presenter presentation.Presenter, repo domain.IRepository) presentation.Controller {
	uploadQuestion := domain.Upload{Repo: repo, Presenter: presenter}
	info := status.Info{Presenter: presenter}
	random := domain.Random{Repo:repo, Presenter:presenter, MaxQuestions:10}
	byId := domain.ById{Repo:repo, Presenter:presenter}

	return presentation.Controller{
		UploadUseCase: uploadQuestion,
		StatusUseCase:info,
		RandomUseCase:random,
		GetById: byId,
	}
}

func (handler Router) Route(router *gin.Engine) {
	q := router.Group("/questions")
	{
		q.POST("/", handler.handle(upload))
		q.GET("/", handler.handle(random))
		q.GET("/:id", handler.handle(questionById))
	}
}

func (handler Router) handle(callback func(presentation.Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		callback(controller(presentation.Presenter{
			Writer: c.Writer,
		}, handler.repo), c)
	}
}

func upload(ctrl presentation.Controller, ctx *gin.Context)       { ctrl.Upload(ctx.Request) }
func random(ctrl presentation.Controller, ctx *gin.Context)       {
	ctrl.Random(ctx.Request)
}
func questionById(ctrl presentation.Controller, ctx *gin.Context) { ctrl.QuestionById(ctx) }
