// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	usertype "ent/types/user"
	"errors"
	"frame/errs"
	"frame/modules/log"
	"frame/modules/resp"
	"service-api/resources/translate/messages"

	"net/http"
	authserver "service-api/internal/app/services/auth"
	"service-api/internal/app/services/captcha"
	"service-api/internal/app/validator/auth"
	"service-api/internal/modules/setting"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignIn return SignIn page before data
func SignIn(ctx *gin.Context) {
	resp.Success(ctx, map[string]any{
		"captcha": setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn),
	})
}

func SignInPost(ctx *gin.Context, form *auth.SignInForm) {
	if setting.CheckCaptchaFeatureEnable(setting.CaptchaFeatureSignIn) {
		st, err := captcha.Verify(form.Captcha.Type, form.Captcha.Token, form.Captcha.Key, form.Captcha.Answer, false)
		if err != nil {
			resp.Error(ctx, err, http.StatusBadRequest, nil)
			return
		}

		if st {
			resp.Error(ctx, errors.New(messages.CaptchaValidationFAILED.ID), http.StatusBadRequest, nil)
			return
		}
	}

	u, source, err := authserver.UserSignIn(ctx, form.UserName, form.Passwd)
	if err != nil {
		if errors.Is(err, errs.ErrNotExist) || errors.Is(err, errs.ErrInvalidArgument) {
			log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Fail(ctx, err.Error(), http.StatusBadRequest, "FORM.USERNAME_PASSWORD_INCORRECT")
		} else if usertype.IsErrUserNotExist(err) {
			log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Fail(ctx, err.Error(), http.StatusBadRequest, "FORM.ACCOUNT_BEEN_USED")
		} else if usertype.IsErrUserProhibitLogin(err) {
			log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Jump(ctx, auth.SignInVerifyError{ProhibitLogin: true}, "SIGN_IN.PROHIBIT_LOGIN")
		} else if usertype.IsErrUserInactive(err) {
			if setting.ServiceSetting.RegisterConfirm {
				log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				resp.SuccessWithMsg(ctx, auth.SignInVerifyError{ActiveYourAccount: true}, "SIGN_IN.ACTIVE_YOUR_ACCOUNT")
			} else {
				log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

				resp.Jump(ctx, auth.SignInVerifyError{ProhibitLogin: true}, "SIGN_IN.PROHIBIT_LOGIN")
			}
		} else {
			log.Sugar().Infof("Failed authentication attempt for %s from %s: %v", form.UserName, ctx.RemoteIP(), err)

			resp.Error(ctx, err, http.StatusInternalServerError, nil)
		}
		return
	}

	// First of all if the source can skip local two fa we're done
	if skipper, ok := source.Cfg.Value().(authserver.LocalTwoFASkipper); ok && skipper.IsSkipLocalTwoFA() {
		if userSession, err := authserver.SignInUserSession(ctx, u, form.Remember); err != nil {
			resp.Error(ctx, err, http.StatusBadRequest, nil)
		} else {
			resp.Success(ctx, userSession)
		}

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
		if userSession, err := authserver.SignInUserSession(ctx, u, form.Remember); err != nil {
			resp.Error(ctx, err, http.StatusBadRequest, nil)
		} else {
			resp.Success(ctx, userSession)
		}
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
