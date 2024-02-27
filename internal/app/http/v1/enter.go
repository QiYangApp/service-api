package v1

type AbstractController struct {
}

type ApiGroup struct {
	CaptchaApi
	AuthApi
}

var Client = new(ApiGroup)
