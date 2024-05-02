// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"ent/models"
	usertype "ent/types/user"
	util "ent/utils"
	"errors"
	"frame/modules/log"
	"frame/modules/resp"

	"net/http"
	"service-api/internal/app/api/validator"
	"service-api/internal/app/api/validator/auth"
	"service-api/internal/modules/setting"
	authserver "service-api/internal/services/auth"
	"service-api/internal/services/captcha"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignIn return SignIn page before data
func SignIn(ctx *gin.Context) {
	resp.Success(ctx, map[string]any{
		"captcha": setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn),
	})
}

func SignInPost(ctx *gin.Context, form auth.SignInForm, captchaVerify *validator.CaptchaVerifyRequest) {
	if setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn) {
		st, err := captcha.Verify(captchaVerify.Type, captchaVerify.Token, captchaVerify.Key, captchaVerify.Answer, false)
		if err != nil {
			resp.Error(ctx, err, http.StatusBadRequest, nil)
			return
		}

		if st {
			resp.Error(ctx, captcha.CaptchaCheckFail, http.StatusBadRequest, nil)
			return
		}
	}

	u, source, err := authserver.UserSignIn(ctx, form.UserName, form.Passwd)
	if err != nil {
		if errors.Is(err, util.ErrNotExist) || errors.Is(err, util.ErrInvalidArgument) {
			log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Fail(ctx, err.Error(), http.StatusBadRequest, "FORM.USERNAME_PASSWORD_INCORRECT")
		} else if usertype.IsErrUserNotExist(err) {
			log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Fail(ctx, err.Error(), http.StatusBadRequest, "FORM.ACCOUNT_BEEN_USED")
		} else if usertype.IsErrUserProhibitLogin(err) {
			log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Jump(ctx, auth.SignInVerifyError{ProhibitLogin: true}, "SIGN_IN.PROHIBIT_LOGIN")
		} else if usertype.IsErrUserInactive(err) {
			if setting.ServiceSetting.RegisterConfirm {
				log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				resp.SuccessWithMsg(ctx, auth.SignInVerifyError{ActiveYourAccount: true}, "SIGN_IN.ACTIVE_YOUR_ACCOUNT")
			} else {
				log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				resp.Jump(ctx, auth.SignInVerifyError{ProhibitLogin: true}, "SIGN_IN.PROHIBIT_LOGIN")
			}
		} else {
			log.Sugar().Info("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Error(ctx, err, http.StatusInternalServerError, nil)
		}
		return
	}

	// First of all if the source can skip local two fa we're done
	if skipper, ok := source.Cfg.Value().(authserver.LocalTwoFASkipper); ok && skipper.IsSkipLocalTwoFA() {
		handleSignIn(ctx, u, form.Remember)
		return
	}

	// No two factor auth configured we can sign in the user
	hasWebAuthnTwofa, err := authserver.HasUserWebAuthn(ctx, u.ID)
	if err != nil {
		log.Sugar().Debug("signIn hasWebAuthnTwofa err", zap.Error(err))
		resp.Error(ctx, err, http.StatusInternalServerError, "")
		return
	}

	hasTOTPtwofa, err := authserver.HasUserTwoFactor(ctx, u.ID)
	if err != nil {
		log.Sugar().Debug("signIn hasTOTPtwofa err", zap.Error(err))
		resp.Error(ctx, err, http.StatusInternalServerError, "")
		return
	}

	if !hasWebAuthnTwofa && !hasTOTPtwofa {
		log.Sugar().Debug("signIn ing, ", form.UserName)
		// No two factor auth configured we can sign in the user
		handleSignIn(ctx, u, form.Remember)
		return
	}

	if hasWebAuthnTwofa {
		resp.Jump(
			ctx,
			auth.SignIn2FAJump{
				WEBAUTHN: true,
			},
			"SIGN_IN.WEBAUTHN",
		)
		return
	}

	resp.Jump(
		ctx,
		auth.SignIn2FAJump{
			TwoFactor: true,
		},
		"SIGN_IN.TWO_FACTOR",
	)
	return
}

func handleSignIn(ctx *gin.Context, u *models.User, remember bool) {

}
