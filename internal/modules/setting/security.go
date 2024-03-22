// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/log"
	"github.com/spf13/viper"
	"service-api/internal/modules/auth/passwd/hash"
)

var SecretSetting = &SecretSettingConfig{}

type SecretSettingConfig struct {
	PasswordComplexity []string `mapstructure:"password_complexity"`
	PasswdHashAlgo     string   `mapstructure:"passwd_hash_algo"`
}

func loadSecret(viper *viper.Viper) {
	if err := viper.Unmarshal(SecretSetting); err != nil {
		log.Client.Sugar().Warnf("load secret setting fail, err: %v", err)
	}

	loadSecretPasswdHashAlgo()
}

func loadSecretPasswdHashAlgo() {
	// Ensure that the provided default hash algorithm is a valid hash algorithm
	var algorithm *hash.PasswordHashAlgorithm
	SecretSetting.PasswdHashAlgo, algorithm = hash.SetDefaultPasswordHashAlgorithm(SecretSetting.PasswdHashAlgo)
	if algorithm == nil {
		log.Client.Sugar().Fatalf("The provided password hash algorithm was invalid: %s", SecretSetting.PasswdHashAlgo)
	}

}
