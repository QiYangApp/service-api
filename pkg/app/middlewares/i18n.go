package middlewares

import (
	"app/helpers"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func I18nUrl() gin.HandlerFunc {
	var getAcceptLanguage func(acceptLanguate string) string
	getAcceptLanguage = func(acceptLanguate string) string {
		var serverLangs = []language.Tag{
			language.SimplifiedChinese, // zh-Hans fallback
			language.Chinese,           // zh-Hans fallback
			language.AmericanEnglish,   // en-US
			language.Korean,            // de
		}

		// 也可以不定义 serverLangs 用下面一行选择支持所有语种。
		// var matcher = language.NewMatcher(message.DefaultCatalog.Languages())
		var matcher = language.NewMatcher(serverLangs)
		t, _, _ := language.ParseAcceptLanguage(acceptLanguate)
		tag, _, _ := matcher.Match(t...)

		//fmt.Printf("best match: %s (%s) index=%d confidence=%v\n",
		//	display.English.Tags().Name(tag),
		//	display.Self.Name(tag),
		//	index, confidence)

		//str := fmt.Sprintf("tag is %s", tag)
		//fmt.Println(str)
		//fmt.Printf("best match: %s\n", display.Self.Name(tag))

		return fmt.Sprintf("%s", tag)
	}

	return func(c *gin.Context) {
		locale := c.Query("locale")
		if locale != "" {
			c.Request.Header.Set("Accept-Language", locale)
		}
		lang := getAcceptLanguage(c.GetHeader("Accept-Language"))

		// NOTE: On June 2012, the deprecation of recommendation to use the "X-" prefix has become official as RFC 6648.
		// https://stackoverflow.com/questions/3561381/custom-http-headers-naming-conventions
		c.Request.Header.Set("I18n-Language", lang)
		c.Next()
	}
}

func I18nLocal() gin.HandlerFunc {
	return i18n.Localize(
		i18n.WithBundle(&i18n.BundleCfg{
			RootPath:         helpers.Path.I18Path,
			AcceptLanguage:   []language.Tag{language.Chinese, language.SimplifiedChinese, language.English},
			DefaultLanguage:  language.SimplifiedChinese,
			UnmarshalFunc:    toml.Unmarshal,
			FormatBundleFile: "toml",
		}),
		//i18n.WithGetLngHandle(
		//	func(r *gin.Context, defaultLng string) string {
		//		/**
		//		**/
		//		lang := r.GetHeader("Accept-Language")
		//		if lang == "" {
		//			return defaultLng
		//		}
		//		return lang
		//	},
		//),
	)
}
