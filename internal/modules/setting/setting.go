// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"frame/conf"
	"frame/modules/log"

	"github.com/spf13/viper"
)

func LoadCommonSettings() {
	if err := loadCommonSettingsFrom(conf.Client()); err != nil {
		log.Sugar().Fatalf("Unable to load settings from config: %v", err)
	}
}

// loadCommonSettingsFrom loads common configurations from a configuration provider.
func loadCommonSettingsFrom(cfg *viper.Viper) error {
	loadApp(cfg)
	loadServiceSetting(cfg.Sub("service"))
	loadTimeFrom(cfg)
	loadSecret(cfg.Sub("security"))
	loadSessionSetting(cfg.Sub("session"))
	loadCaptchaSetting(cfg.Sub("captcha"))
	loadAuthSetting(cfg.Sub("auth"))

	return nil
}
