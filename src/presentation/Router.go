package presentation

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type IRouter interface {
	Route(router *gin.RouterGroup)
}

func NewRouter(topics IRouter, essays IRouter, status IRouter, questions IRouter, tag IRouter, port int) {
	router := gin.Default()
	status.Route(router.Group("/"))
	questions.Route(router.Group("/questions"))
	essays.Route(router.Group("/essay"))
	topics.Route(router.Group("/topic"))
	tag.Route(router.Group("/tag"))
	err := router.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}