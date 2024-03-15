// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/log"
	"github.com/spf13/viper"
	"math"
	"service-api/internal/modules/generate"
	"sync/atomic"
)

// OAuth2UsernameType is enum describing the way gitea 'name' should be generated from oauth2 data
type OAuth2UsernameType string

const (
	// OAuth2UsernameUserid oauth2 userid field will be used as gitea name
	OAuth2UsernameUserid OAuth2UsernameType = "userid"
	// OAuth2UsernameNickname oauth2 nickname field will be used as gitea name
	OAuth2UsernameNickname OAuth2UsernameType = "nickname"
	// OAuth2UsernameEmail username of oauth2 email field will be used as gitea name
	OAuth2UsernameEmail OAuth2UsernameType = "email"
)

func (username OAuth2UsernameType) isValid() bool {
	switch username {
	case OAuth2UsernameUserid, OAuth2UsernameNickname, OAuth2UsernameEmail:
		return true
	}
	return false
}

// OAuth2AccountLinkingType is enum describing behaviour of linking with existing account
type OAuth2AccountLinkingType string

const (
	// OAuth2AccountLinkingDisabled error will be displayed if account exist
	OAuth2AccountLinkingDisabled OAuth2AccountLinkingType = "disabled"
	// OAuth2AccountLinkingLogin account linking login will be displayed if account exist
	OAuth2AccountLinkingLogin OAuth2AccountLinkingType = "login"
	// OAuth2AccountLinkingAuto account will be automatically linked if account exist
	OAuth2AccountLinkingAuto OAuth2AccountLinkingType = "auto"
)

func (accountLinking OAuth2AccountLinkingType) isValid() bool {
	switch accountLinking {
	case OAuth2AccountLinkingDisabled, OAuth2AccountLinkingLogin, OAuth2AccountLinkingAuto:
		return true
	}
	return false
}

// OAuth2Client settings
var OAuth2Client struct {
	RegisterEmailConfirm   bool
	OpenIDConnectScopes    []string
	EnableAutoRegistration bool
	Username               OAuth2UsernameType
	UpdateAvatar           bool
	AccountLinking         OAuth2AccountLinkingType
}

func loadOAuth2ClientFrom(cfg *viper.Viper) {
	OAuth2Client.RegisterEmailConfirm = cfg.GetBool("oauth2_client.register_email_confirm")
	OAuth2Client.OpenIDConnectScopes = parseScopes(cfg, "oauth2_client.openid_connect_scopes")
	OAuth2Client.EnableAutoRegistration = cfg.GetBool("oauth2_client.enable_auto_registration")
	OAuth2Client.Username = OAuth2UsernameType(cfg.GetString("USERNAME"))
	if !OAuth2Client.Username.isValid() {
		log.Client().Sugar().Warn("Username setting is not valid: '%s', will fallback to '%s'", OAuth2Client.Username, OAuth2UsernameNickname)
		OAuth2Client.Username = OAuth2UsernameNickname
	}
	OAuth2Client.UpdateAvatar = cfg.GetBool("oauth2_client.update_avatar")
	OAuth2Client.AccountLinking = OAuth2AccountLinkingType(cfg.GetString("oauth2_client.account_linking"))
	if !OAuth2Client.AccountLinking.isValid() {
		log.Client().Sugar().Warn("Account linking setting is not valid: '%s', will fallback to '%s'", OAuth2Client.AccountLinking, OAuth2AccountLinkingLogin)
		OAuth2Client.AccountLinking = OAuth2AccountLinkingLogin
	}
}

func parseScopes(cfg *viper.Viper, name string) []string {
	parts := cfg.GetStringMapString(name)
	scopes := make([]string, 0, len(parts))
	for _, scope := range parts {
		if scope != "" {
			scopes = append(scopes, scope)
		}
	}
	return scopes
}

var OAuth2 = struct {
	Enabled                    bool
	AccessTokenExpirationTime  int64
	RefreshTokenExpirationTime int64
	InvalidateRefreshTokens    bool
	JWTSigningAlgorithm        string `ini:"JWT_SIGNING_ALGORITHM"`
	JWTSigningPrivateKeyFile   string `ini:"JWT_SIGNING_PRIVATE_KEY_FILE"`
	MaxTokenLength             int
	DefaultApplications        []string
}{
	Enabled:                    true,
	AccessTokenExpirationTime:  3600,
	RefreshTokenExpirationTime: 730,
	InvalidateRefreshTokens:    false,
	JWTSigningAlgorithm:        "RS256",
	JWTSigningPrivateKeyFile:   "jwt/private.pem",
	MaxTokenLength:             math.MaxInt16,
	DefaultApplications:        []string{"git-credential-oauth", "git-credential-manager", "tea"},
}

func loadOAuth2From(cfg *viper.Viper) {
	//if err := sec.MapTo(&OAuth2); err != nil {
	//	log.Client().Sugar().Fatal("Failed to map OAuth2 settings: %v", err)
	//	return
	//}
	//
	//// Handle the rename of ENABLE to ENABLED
	//deprecatedSetting(rootCfg, "oauth2", "ENABLE", "oauth2", "ENABLED", "v1.23.0")
	//if sec.HasKey("ENABLE") && !sec.HasKey("ENABLED") {
	//	OAuth2.Enabled = sec.Key("ENABLE").MustBool(OAuth2.Enabled)
	//}
	//
	//if !OAuth2.Enabled {
	//	return
	//}
	//
	//jwtSecretBase64 := loadSecret(sec, "JWT_SECRET_URI", "JWT_SECRET")
	//
	//if !filepath.IsAbs(OAuth2.JWTSigningPrivateKeyFile) {
	//	OAuth2.JWTSigningPrivateKeyFile = filepath.Join(AppDataPath, OAuth2.JWTSigningPrivateKeyFile)
	//}
	//
	//if InstallLock {
	//	jwtSecretBytes, err := generate.DecodeJwtSecretBase64(jwtSecretBase64)
	//	if err != nil {
	//		jwtSecretBytes, jwtSecretBase64, err = generate.NewJwtSecretWithBase64()
	//		if err != nil {
	//			log.Client().Sugar().Fatal("error generating JWT secret: %v", err)
	//		}
	//		saveCfg, err := rootCfg.PrepareSaving()
	//		if err != nil {
	//			log.Client().Sugar().Fatal("save oauth2.JWT_SECRET failed: %v", err)
	//		}
	//		rootCfg.Section("oauth2").Key("JWT_SECRET").SetValue(jwtSecretBase64)
	//		saveCfg.Section("oauth2").Key("JWT_SECRET").SetValue(jwtSecretBase64)
	//		if err := saveCfg.Save(); err != nil {
	//			log.Client().Sugar().Fatal("save oauth2.JWT_SECRET failed: %v", err)
	//		}
	//	}
	//	generalSigningSecret.Store(&jwtSecretBytes)
	//}
}

// generalSigningSecret is used as container for a []byte value
// instead of an additional mutex, we use CompareAndSwap func to change the value thread save
var generalSigningSecret atomic.Pointer[[]byte]

func GetGeneralTokenSigningSecret() []byte {
	old := generalSigningSecret.Load()
	if old == nil || len(*old) == 0 {
		jwtSecret, _, err := generate.NewJwtSecretWithBase64()
		if err != nil {
			log.Client().Sugar().Fatal("Unable to generate general JWT secret: %s", err.Error())
		}
		if generalSigningSecret.CompareAndSwap(old, &jwtSecret) {
			// FIXME: in main branch, the signing token should be refactored (eg: one unique for LFS/OAuth2/etc ...)
			log.Client().Sugar().Warn("OAuth2 is not enabled, unable to use a persistent signing secret, a new one is generated, which is not persistent between restarts and cluster nodes")
			return jwtSecret
		}
		return *generalSigningSecret.Load()
	}
	return *old
}
