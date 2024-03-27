package api

import v1 "service-api/internal/app/api/v1"

type AbstractController struct {
}

type Group struct {
	v1.CaptchaApi
}

var Client = new(Group)
