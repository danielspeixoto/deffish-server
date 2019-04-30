package tag

import (
	qboundary "deffish-server/src/boundary/question"
	boundary "deffish-server/src/boundary/tag"
	"deffish-server/src/domain/tag"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ctrl Controller
}

func NewRouter(r boundary.IRepository, q qboundary.IRepository) Router {
	ctrl := Controller{
		SuggestionsUseCase: tag.SuggestionsUseCase{r, q},
		ByNameUseCase:      tag.ByNameUseCase{r},
		UploadUseCase:      tag.UploadUseCase{r},
	}
	return Router{
		ctrl,
	}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.POST("/", r.ctrl.Post)
	router.GET("/", r.ctrl.Get)
}