package system

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type LanguageService struct {
	service
}

func (d *LanguageService) Handle(r *gin.Engine, cfg *ConfigService) {

	r.Use(i18n.Localize(
		i18n.WithBundle(&i18n.BundleCfg{
			RootPath:         "./src/resources/lang",
			AcceptLanguage:   []language.Tag{language.SimplifiedChinese, language.English},
			DefaultLanguage:  language.SimplifiedChinese,
			UnmarshalFunc:    toml.Unmarshal,
			FormatBundleFile: "toml",
		}),
		i18n.WithGetLngHandle(
			func(r *gin.Context, defaultLng string) string {
				/**
				**/
				lang := r.GetHeader("Accept-Language")
				if lang == "" {
					return defaultLng
				}
				return lang
			},
		),
	))
}
