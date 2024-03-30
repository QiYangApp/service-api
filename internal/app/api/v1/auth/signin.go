// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	usertype "ent/types/user"
	util "ent/utils"
	"errors"
	"framework/log"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/internal/app/api/validator"
	"service-api/internal/app/api/validator/auth"
	"service-api/internal/modules/setting"
	authserver "service-api/internal/services/auth"
	"service-api/internal/services/captcha"
)

// SignIn return SignIn page before data
func SignIn(ctx *gin.Context, a validator.Test) {
	response.RSuccess(ctx, map[string]any{
		"captcha": setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn),
	})
}

func SignInPost(ctx *gin.Context, form auth.SignInForm, captchaVerify *validator.CaptchaVerifyRequest) {
	if setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn) {
		st, err := captcha.Verify(captchaVerify.Type, captchaVerify.Token, captchaVerify.Key, captchaVerify.Answer, false)
		if err != nil {
			response.RError(ctx, err, http.StatusBadRequest, nil)
			return
		}

		if st {
			response.RError(ctx, captcha.CaptchaCheckFail, http.StatusBadRequest, nil)
			return
		}
	}

	_, _, err := authserver.UserSignIn(ctx, form.UserName, form.Passwd)
	if err != nil {
		if errors.Is(err, util.ErrNotExist) || errors.Is(err, util.ErrInvalidArgument) {
			log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			response.RFail(ctx, err.Error(), http.StatusBadRequest, "FORM.USERNAME_PASSWORD_INCORRECT")
		} else if usertype.IsErrUserNotExist(err) {
			log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			response.RFail(ctx, err.Error(), http.StatusBadRequest, "FORM.ACCOUNT_BEEN_USED")
		} else if usertype.IsErrUserProhibitLogin(err) {
			log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			response.RJump(ctx, map[string]any{
				"PROHIBIT_LOGIN": true,
			}, "FORM.PROHIBIT_LOGIN")
		} else if usertype.IsErrUserInactive(err) {
			if setting.ServiceSetting.RegisterConfirm {
				log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				response.RSuccessWithMsg(ctx, map[string]any{
					"ACTIVE_YOUR_ACCOUNT": true,
				}, "FORM.PROHIBIT_LOGIN")
			} else {
				log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				response.RJump(ctx, map[string]any{
					"PROHIBIT_LOGIN": true,
				}, "FORM.PROHIBIT_LOGIN")
			}
		} else {
			log.Client.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			response.RError(ctx, err, http.StatusInternalServerError, nil)
		}
		return
	}

	// First of all if the source can skip local two fa we're done
	//if skipper, ok := source.Cfg.(auth_service.LocalTwoFASkipper); ok && skipper.IsSkipLocalTwoFA() {
	//	handleSignIn(ctx, u, form.Remember)
	//	return
	//}
}

func handleSignIn(ctx *gin.Context) {

}
