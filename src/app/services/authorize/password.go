package authorize

import "service-api/src/app/entity/http"

type PasswordLoginService[P http.VerifyType] struct {
	AuthorizedService[P]
}

func (s *PasswordLoginService[P]) Authorizing(p P) any {
	return nil
}

func (s *PasswordLoginService[P]) Authorized(p P) any {

	return nil
}
