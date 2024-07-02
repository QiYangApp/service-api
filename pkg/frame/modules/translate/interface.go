package translate

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// GinI18n ...
type GinI18n interface {
	getDefaultLanguage(context *gin.Context) language.Tag
	getMessage(context *gin.Context, param interface{}) (string, error)
	mustGetMessage(context *gin.Context, param interface{}) string
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
	hasLang(language string) bool
}
