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
	handler.handle("/status", status)
	handler.handle("/questions/", upload)
	handler.handle("/questions/random", random)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func (handler Handler) handle(pattern string, callback func(Controller, *http.Request)) {
	http.HandleFunc(pattern,
		func(writer http.ResponseWriter, request *http.Request) {

			log.Printf("Request at %s from IP: %s", request.URL.Path, request.RemoteAddr)

			ctrl := controller(Presenter{Writer: writer}, handler.Repo)
			callback(ctrl, request)
		})
}

func controller(presenter Presenter, repo gateway.IQuestionRepository) Controller {
	uploadQuestion := usecase.UploadQuestionUseCase{Repo: repo, Presenter: presenter}
	status := usecase.StatusUseCase{Repo: repo, Presenter: presenter}
	random := usecase.RandomQuestionUseCase{Repo:repo, Presenter:presenter, MaxQuestions:10}
	return Controller{UploadQuestionUseCase: uploadQuestion,
		StatusUseCase:status, RandomQuestionUseCase:random}
}

func upload(c Controller, r *http.Request) { c.Upload(r) }
func status(c Controller, r *http.Request) { c.Status(r) }
func random(c Controller, r *http.Request) { c.Random(r) }
