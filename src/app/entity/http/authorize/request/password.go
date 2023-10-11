package request

import "service-api/src/app/entity/http"

type PasswordAccountCheckVerify struct {
	http.VerifyType
	Account string `form:"account"`
}

type PasswordLoginVerify struct {
	http.VerifyType
	Account string `form:"account"`
}
