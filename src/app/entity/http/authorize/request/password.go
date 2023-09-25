package request

import "service-api/src/app/entity/http"

type PasswordLoginVerify struct {
	http.VerifyType
	Account string `form:"account"`
}
