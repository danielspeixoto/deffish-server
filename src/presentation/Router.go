package presentation

import "C"
import (
	"deffish-server/src/domain/gateway"
	"deffish-server/src/domain/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	Repo gateway.IQuestionRepository
}

func NewHandler(repo gateway.IQuestionRepository, port int) {
	handler := Handler{
		repo,
	}
	router := gin.Default()

	router.GET("/", handler.handle(status))
	router.GET("/status", handler.handle(status))
	router.POST("/questions", handler.handle(upload))
	router.GET("/questions", handler.handle(random))
	router.GET("/questions/:id", handler.handle(questionById))

	err := router.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}

func (h Handler) handle(callback func(Controller, *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		ctrl := controller(Presenter{Writer: c.Writer}, h.Repo)
		callback(ctrl, c)
	}
}

func controller(presenter Presenter, repo gateway.IQuestionRepository) Controller {
	uploadQuestion := usecase.UploadQuestionUseCase{Repo: repo, Presenter: presenter}
	status := usecase.StatusUseCase{Repo: repo, Presenter: presenter}
	random := usecase.RandomQuestionUseCase{Repo:repo, Presenter:presenter, MaxQuestions:10}
	questionById := usecase.QuestionByIdUseCase{Repo:repo, Presenter:presenter}
	return Controller{UploadQuestionUseCase: uploadQuestion,
		StatusUseCase:status, RandomQuestionUseCase:random,
		GetById: questionById,
	}
}

func upload(ctrl Controller, ctx *gin.Context)       { ctrl.Upload(ctx.Request) }
func status(ctrl Controller, ctx *gin.Context)       { ctrl.Status(ctx.Request) }
func random(ctrl Controller, ctx *gin.Context)       { ctrl.Random(ctx.Request) }
func questionById(ctrl Controller, ctx *gin.Context) { ctrl.QuestionById(ctx) }
