package routing

import "C"
import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type IRouter interface {
	Route(router *gin.RouterGroup)
}

func NewRouter(status IRouter, questions IRouter, port int) {
	router := gin.Default()
	status.Route(router.Group("/"))
	questions.Route(router.Group("/question"))
	err := router.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}