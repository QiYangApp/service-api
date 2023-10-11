package authorize

import (
	"service-api/src/app/entity/http"
)

type AuthorizedService[P http.VerifyType] interface {
	Check(p P) interface{}

	Authorizing(p P) interface{}

	Authorized(p P) interface{}
}
