package request

import "service-api/src/app/entity/http"

type PasswordLoginVerify struct {
	http.VerifyType
	account string `form:"account"`
}
