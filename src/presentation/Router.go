package presentation

import (
	"deffish-server/src/domain/gateway"
	"deffish-server/src/domain/usecase"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Repo gateway.IQuestionRepository
}

func NewHandler(repo gateway.IQuestionRepository, port int) {
	handler := Handler{
		repo,
	}
	http.HandleFunc("/upload", handler.upload)
	http.HandleFunc("/status", handler.status)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}


func (handler Handler) upload(writer http.ResponseWriter, request *http.Request) {
	ctrl := controller(Presenter{Writer: writer}, handler.Repo)
	ctrl.Upload(request)
}

func (handler Handler) status(writer http.ResponseWriter, request *http.Request) {
	ctrl := controller(Presenter{Writer: writer}, handler.Repo)
	ctrl.Status(request)
}

func controller(presenter Presenter, repo gateway.IQuestionRepository) Controller {
	uploadQuestion := usecase.UploadQuestionUseCase{Repo: repo, Presenter: presenter}
	status := usecase.StatusUseCase{Repo: repo, Presenter: presenter}
	return Controller{UploadQuestionUseCase: uploadQuestion, StatusUseCase:status}
}
