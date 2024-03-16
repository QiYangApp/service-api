package validator

type PreAuthDataRequest struct {
	Type string `uri:"type"`
}

type PreAuthVerifyRequest struct {
	Type    string `uri:"type"`
	Account string `form:"account" query:"account"`
}

type AuthorizingRequest struct {
	Type string `uri:"type"`
	Body any    `form:"body"`
}

type AuthorizedRequest struct {
	Type string `uri:"type"`
	Body any    `form:"body"`
}
