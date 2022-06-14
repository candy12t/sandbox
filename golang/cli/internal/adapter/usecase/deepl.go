package usecase

type DeepL interface {
	Translate(original, sourceLanguage, targetLanguage string) (string, error)
}
