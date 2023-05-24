package login

import (
	"service-api/src/api/controller"
	"service-api/src/app/entity/http/authorize/request"
	"service-api/src/core/helpers/response"
)

type PasswordController struct {
	controller.AbstractController
}

func (PasswordController) Edit(p request.PasswordLoginVerify) *response.Response[bool] {
	return response.Single[bool](true)
}
