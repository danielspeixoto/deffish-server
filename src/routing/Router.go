package routing

import "C"
import (
	"deffish-server/src/question"
	"github.com/gin-gonic/gin"
	"strconv"
)



func NewRouter(questions question.Router, port int) {
	router := gin.Default()
	questions.Route(router)
	err := router.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}