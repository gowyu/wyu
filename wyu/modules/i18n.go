package modules

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"strings"
)

const (
	dirLanguage string = "./resources/lang/"
)

func I18nT(key string, ln string) string {
	Tag, _ := Translated.LanguageFormat(ln)
	return translations[Tag].Sprintf(key)
}

type I18N interface {
	Loading() error
	LanguageFormat(ln string) (language.Tag, error)
}

type i18n struct {

}

var (
	_ I18N = &i18n{}

	Translated I18N
	translations map[language.Tag]*message.Printer
)

func NewI18N() *i18n {
	return &i18n{}
}

func (translate *i18n) Loading() error {
	dir := Env.GET("Languages.Dir", dirLanguage).(string)

	f, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	translations = make(map[language.Tag]*message.Printer, 0)
	for _, val := range f {
		if val.IsDir() {
			continue
		}

		fn := strings.Split(val.Name(), ".")
		ln := fn[0]

		if fn[1] != "json" {
			continue
		}

		byteJson, errTranslated := ioutil.ReadFile(dir + val.Name())
		if errTranslated != nil {
			return errTranslated
		}

		for key, translated := range UtilsJsonToMap(byteJson) {
			message.SetString(language.MustParse(ln), key, translated.(string))
		}

		Tag, errTag := translate.LanguageFormat(ln)
		if errTag != nil {
			panic(errTag.Error())
		}

		translations[Tag] = message.NewPrinter(Tag)
	}

	Translated = translate
	return nil
}

func (translate *i18n) LanguageFormat(ln string) (language.Tag, error) {
	if ok, _ := UtilsStrContains(ln, "cn"); ok {
		return language.Chinese, nil
	}

	languages := Env.GET("Languages.Lns", []interface{}{}).([]interface{})
	if ok, _ := UtilsStrContains(ln, languages ...); ok == false {
		return language.English, errors.New("lost env languages configs")
	}

	return language.MustParse(ln), nil
}