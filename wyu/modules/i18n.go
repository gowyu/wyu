package modules

import (
	"errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"strings"
)

const dirLanguage string = "./resources/lang/"

func I18nT(key string, ln string) string {
	tag, _ := Translated.LanguageFormat(ln)
	return translations[tag].Sprintf(key)
}

type i18n struct {}

var (
	Translated *i18n
	translations map[language.Tag]*message.Printer
)

func NewI18N() *i18n {
	return &i18n{}
}

func (translate *i18n) Loading() (err error) {
	dir := Env.GET("Languages.Dir", dirLanguage).(string)

	f, err := ioutil.ReadDir(dir)
	if err != nil {
		return
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

		byteJson, err := ioutil.ReadFile(dir + val.Name())
		if err != nil {
			return err
		}

		for key, translated := range UtilsJsonToMap(byteJson) {
			message.SetString(language.MustParse(ln), key, translated.(string))
		}

		Tag, err := translate.LanguageFormat(ln)
		if err != nil {
			return err
		}

		translations[Tag] = message.NewPrinter(Tag)
	}

	Translated = translate
	return
}

func (translate *i18n) LanguageFormat(ln string) (tag language.Tag, err error) {
	if ok, _ := UtilsStrContains(ln, "cn"); ok {
		tag = language.Chinese
		return
	}

	languages := Env.GET("Languages.Lns", []interface{}{}).([]interface{})
	if ok, _ := UtilsStrContains(ln, languages ...); ok == false {
		tag = language.English
		err = errors.New("lost env languages configs")
		return
	}

	tag = language.MustParse(ln)
	return
}