package scripts

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var translator = i18n.NewLocalizer(
	i18n.NewBundle(language.English),
	language.SimplifiedChinese.String(),
	language.Chinese.String(),
)

func main() {
	//localizer.Localize(&i18n.LocalizeConfig{
	//	DefaultMessage: &i18n.Message{
	//		ID:    "PersonCats",
	//		One:   "{{.Name}} has {{.Count}} cat.",
	//		Other: "{{.Name}} has {{.Count}} cats.",
	//	},
	//	TemplateData: map[string]interface{}{
	//		"Name":  "Nick",
	//		"Count": 2,
	//	},
	//	PluralCount: 2,
	//}) // Nick has 2 cats.

	translator.LocalizeMessage(
		&i18n.Message{
			ID:    "WELCOME",
			Other: "Welcome !",
		},
	)
}
