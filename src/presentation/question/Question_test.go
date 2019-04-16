package question

import "deffish-server/src/aggregates"

var exampleQuestion = aggregates.Question{
	Image: aggregates.Image{},
	Source: "ENEM",
	Variant: "Amarelo",
	Edition: 2017,
	Number: 3,
	Domain: "Linguagens",
	Answer: 1,
}