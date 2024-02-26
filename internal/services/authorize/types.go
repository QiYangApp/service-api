package authorize

import "service-api/internal/entity"

type Authorizing interface {
	PreAuthorizationData(req entity.LoginRequest) *entity.LoginResponse

	PreAuthorizationVerification(req entity.LoginRequest) *entity.LoginResponse

	Authorizing(req entity.LoginRequest) *entity.LoginResponse

	Authorized(req entity.LoginRequest) *entity.LoginResponse
}
