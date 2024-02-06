package controller

import (
	"service-api/src/service/authorize"
)

type PasswordLogin struct {
	Service authorize.PasswordLoginService
}

func (*PasswordLogin) Login() {

}
