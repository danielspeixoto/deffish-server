package presentation

import (
	"deffish-server/src/status"
	"net/http"
)

type Controller struct {
	StatusUseCase         status.IInfoUseCase
}

func (ctrl Controller) Status(request *http.Request) {
	ctrl.StatusUseCase.Info()
}