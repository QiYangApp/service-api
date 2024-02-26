package validator

import "github.com/google/uuid"

type PasswordSingUpVerificationRequest struct {
	Account  string `form:"account" json:"account" binding:"account" `
	Password string `form:"password" json:"password" binding:"password" `
	Code     string `form:"code" json:"code" binding:"code" `
}

type LoginResponse struct {
	UserId uuid.UUID
	Token  string
}
