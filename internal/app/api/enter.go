package api

import v1 "service-api/internal/app/api/v1"

type AbstractController struct {
}

type Group struct {
	v1.CaptchaApi
	v1.AuthApi
}

var Client = new(Group)
