package request

import (
	"service-api/src/app/entity/http"
)

type PasswordLoginCheckReq struct {
	http.ReqType `json:"_"`
	Account      string `form:"account"`
}

type PasswordLoginCheckRsp struct {
	http.RespType `json:"_"`
	Account       string `form:"account"`
}

type PasswordLoggingReq struct {
	http.ReqType
	Account string `form:"account"`
}

type PasswordLoggingRsp struct {
	http.RespType
	Account string `form:"account"`
}

type PasswordLoggedReq struct {
	http.ReqType
	Account string `form:"account"`
}

type PasswordLoggedRsp struct {
	http.RespType
	Account string `form:"account"`
}
