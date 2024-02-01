package exceptions

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//https://juejin.cn/post/7015517416608235534

// Validator 验证器接口
type Validator interface {
	// GetMessage 获取验证器自定义错误信息
	GetMessage(c *gin.Context) ValidatorMessages
}

// ValidatorMessages 验证器自定义错误信息字典
type ValidatorMessages map[string]string

// GetErrorMsg 获取自定义错误信息
func GetErrorMsg(c *gin.Context, request Validator, err error) map[string]string {
	var msg = make(map[string]string)
	for _, v := range err.(validator.ValidationErrors) {
		key := v.Field() + "." + v.Tag()

		msg[key] = v.Error()
		if message, exist := request.GetMessage(c)[key]; exist {
			msg[key] = message
		}
	}
	return msg
}
