package controller

import "service-api/api/http/controller/authorize"

type AbstractController struct {
}

type ApiGroup struct {
	AuthorizeApi authorize.GroupApi
	CaptchaApi
}

var Client = new(ApiGroup)
