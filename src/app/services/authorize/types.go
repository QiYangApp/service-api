package authorize

import (
	"service-api/src/app/entity/http"
)

type AuthorizedService interface {
	Check(p http.ReqType) http.RespType

	Authorizing(p http.ReqType) http.RespType

	Authorized(p http.ReqType) http.RespType
}
