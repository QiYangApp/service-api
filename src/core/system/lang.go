package system

import (
	"service-api/src/core/config"
	"service-api/src/core/helpers"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type LanguageService struct {
	service
}

func (d *LanguageService) Handle(r *gin.Engine, cfg *config.ConfigService) {

	r.Use(i18n.Localize(
		i18n.WithBundle(&i18n.BundleCfg{
			RootPath:         helpers.PathInstance.JoinCurrentRunPath("resources/lang"),
			AcceptLanguage:   []language.Tag{language.Chinese, language.SimplifiedChinese, language.English},
			DefaultLanguage:  language.Chinese,
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
