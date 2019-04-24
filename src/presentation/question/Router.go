package question

import (
	questionBoundary "deffish-server/src/boundary/question"
	relatedVideoBoundary "deffish-server/src/boundary/relatedVideo"
	questionDomain "deffish-server/src/domain/question"
	relatedVideoDomain "deffish-server/src/domain/relatedVideo"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ctrl Controller
}

func NewRouter(questionRepo questionBoundary.IRepository, videoRepo relatedVideoBoundary.IRepository) Router {
	ctrl := Controller{
		UploadUseCase: questionDomain.Upload{questionRepo},
		RandomTagsUseCase: questionDomain.RandomByTags{questionRepo, 10},
		GetById: questionDomain.ById{questionRepo},
		Videos: relatedVideoDomain.FilterByQuestionUseCase{videoRepo},
	}
	return Router{
		ctrl,
	}
}

func (handler Router) Route(router *gin.RouterGroup) {
	router.POST("/", handler.ctrl.Upload)
	router.GET("/", handler.ctrl.RandomByTags)
	router.GET("/:id/relatedVideos", handler.ctrl.RelatedVideos)
	router.GET("/:id/", handler.ctrl.QuestionById)
}