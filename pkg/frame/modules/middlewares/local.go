package middlewares

import (
	"fmt"
	"frame/util/path"
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
		c.Request.Header.Set("Local-Language", lang)
		c.Next()
	}
}

func I18nLocal() gin.HandlerFunc {
	return i18n.Localize(
		i18n.WithBundle(&i18n.BundleCfg{
			RootPath:         path.RunI18nPath,
			AcceptLanguage:   []language.Tag{language.Chinese, language.SimplifiedChinese, language.English},
			DefaultLanguage:  language.SimplifiedChinese,
			UnmarshalFunc:    toml.Unmarshal,
			FormatBundleFile: "toml",
		}),
		i18n.WithGetLngHandle(
			func(r *gin.Context, defaultLng string) string {
				// 1. Check URL arguments.
				lang := r.Query("lang")
				changeLang := lang != ""

				// 2. Get language information from cookies.
				if len(lang) == 0 {
					ck, _ := r.Cookie("lang")
					if ck != "" {
						lang = ck
					}
				}

				// Check again in case someone changes the supported language list.
				if lang != "" && !i18n {
					lang = ""
					changeLang = false
				}

				lang := r.GetHeader("Accept-Language")
				if lang == "" {
					return defaultLng
				}
				return lang
			},
		),
	)
}

// SetLocaleCookie convenience function to set the locale cookie consistently
func SetLocaleCookie(c *gin.Context, lang string, maxAge int) {
	SetSiteCookie(c, "lang", lang, maxAge)
}

// DeleteLocaleCookie convenience function to delete the locale cookie consistently
// Setting the lang cookie will trigger the middleware to reset the language to previous state.
func DeleteLocaleCookie(c *gin.Context) {
	SetSiteCookie(c, "lang", "", -1)
}
