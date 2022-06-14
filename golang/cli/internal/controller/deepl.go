package controller

import (
	"fmt"
	"log"

	"github.com/candy12t/cli/internal/usecase"
	"github.com/candy12t/cli/internal/webapi"
)

type Form struct {
	sourceLanguage string
	targetLanguage string
}

func (f *Form) Translate(original string) {
	api := webapi.New()
	us := usecase.New(api)
	text, err := us.Translate(original, f.sourceLanguage, f.targetLanguage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
