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
	http.HandleFunc("/status", handler.controllerCall(status))
	http.HandleFunc("/questions/random", handler.controllerCall(random))
	http.HandleFunc("questions/", handler.controllerCall(upload))

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func (handler Handler) controllerCall(callback func(Controller, *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Request at %s", request.URL.Path)
		ctrl := controller(Presenter{Writer: writer}, handler.Repo)
		callback(ctrl, request)
	}
}

func controller(presenter Presenter, repo gateway.IQuestionRepository) Controller {
	uploadQuestion := usecase.UploadQuestionUseCase{Repo: repo, Presenter: presenter}
	status := usecase.StatusUseCase{Repo: repo, Presenter: presenter}
	random := usecase.RandomQuestionUseCase{Repo:repo, Presenter:presenter, MaxQuestions:10}
	return Controller{UploadQuestionUseCase: uploadQuestion,
		StatusUseCase:status, RandomQuestionUseCase:random}
}

func upload(ctrl Controller, request *http.Request) { ctrl.Upload(request) }
func status(ctrl Controller, request *http.Request) { ctrl.Status(request) }
func random(ctrl Controller, request *http.Request) { ctrl.Random(request) }
