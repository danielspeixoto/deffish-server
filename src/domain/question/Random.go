package question

import (
	"log"
)

type RandomByDomain struct {
	Repo IRepository
	Presenter IRandomPresenter
	MaxQuestions int
}

var _ IRandomByDomainUseCase = (*RandomByDomain)(nil)

func (useCase RandomByDomain) Random(amount int, domain string) {
	if amount > useCase.MaxQuestions {
		log.Printf("An user requested %v questions, max is %v", amount, useCase.MaxQuestions)
		amount = useCase.MaxQuestions
	}
	questions, err := useCase.Repo.RandomByDomain(amount, domain)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(questions)
	}
}

type RandomByTags struct {
	Repo IRepository
	Presenter IRandomPresenter
	MaxQuestions int
}

var _ IRandomByTagsUseCase = (*RandomByTags)(nil)

func (useCase RandomByTags) Random(amount int, tags []string) {
	if amount > useCase.MaxQuestions {
		log.Printf("An user requested %v questions, max is %v", amount, useCase.MaxQuestions)
		amount = useCase.MaxQuestions
	}
	questions, err := useCase.Repo.RandomByTags(amount, tags)
	if err != nil {
		log.Print(err)
		useCase.Presenter.OnError(err)
	} else {
		useCase.Presenter.OnListReceived(questions)
	}
}
