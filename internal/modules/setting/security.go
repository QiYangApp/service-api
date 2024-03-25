// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/log"
	"github.com/spf13/viper"
	"service-api/internal/modules/auth/passwd/hash"
)

var SecretSettingConfig = &SecretSetting{}

type SecretSetting struct {
	PasswordComplexity []string `mapstructure:"password_complexity"`
	PasswdHashAlgo     string   `mapstructure:"passwd_hash_algo"`
	PasswdCheckPwn     bool     `mapstructure:"passwd_check_pwn"`
}

func loadSecret(viper *viper.Viper) {
	if err := viper.Unmarshal(SecretSettingConfig); err != nil {
		log.Client.Sugar().Warnf("load secret setting fail, err: %v", err)
	}

	loadSecretPasswdHashAlgo()
}

func loadSecretPasswdHashAlgo() {
	// Ensure that the provided default hash algorithm is a valid hash algorithm
	var algorithm *hash.PasswordHashAlgorithm
	SecretSettingConfig.PasswdHashAlgo, algorithm = hash.SetDefaultPasswordHashAlgorithm(SecretSettingConfig.PasswdHashAlgo)
	if algorithm == nil {
		log.Client.Sugar().Fatalf("The provided password hash algorithm was invalid: %s", SecretSettingConfig.PasswdHashAlgo)
	}

}
