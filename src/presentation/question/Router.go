package question

import (
	boundary "deffish-server/src/boundary/question"
	"deffish-server/src/boundary/tag"
	domain "deffish-server/src/domain/question"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ctrl Controller
}

func NewRouter(questionRepo boundary.IRepository, tagRepo tag.IRepository) Router {
	ctrl := Controller{
		UploadUseCase:     domain.Upload{questionRepo},
		RandomTagsUseCase: domain.RandomByTags{questionRepo, 10},
		GetByIdUseCase:    domain.ById{questionRepo},
		VideosUseCase:     domain.FilterByQuestionUseCase{questionRepo},
		AddTagUseCase:     domain.AddTag{questionRepo, tagRepo},
	}
	return Router{
		ctrl,
	}
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.ctrl.Upload)
	router.GET("/", handler.ctrl.RandomByTags)
	router.GET("/:id/relatedVideos", handler.ctrl.RelatedVideos)
	router.POST("/:id/tags", handler.ctrl.AddTag)
	router.GET("/:id/", handler.ctrl.QuestionById)
}