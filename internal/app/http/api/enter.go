package api

import "service-api/internal/app/http/api/authorize"

type AbstractController struct {
}

type ApiGroup struct {
	AuthorizeApi authorize.GroupApi
	CaptchaApi
}

var Client = new(ApiGroup)
