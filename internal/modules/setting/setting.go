// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2017 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/config"
	"github.com/spf13/viper"
	"log"
)

func LoadCommonSettings() {
	if err := loadCommonSettingsFrom(config.Client); err != nil {
		log.Fatal("Unable to load settings from config: %v", err)
	}
}

// loadCommonSettingsFrom loads common configurations from a configuration provider.
func loadCommonSettingsFrom(cfg *viper.Viper) error {
	loadApp(cfg)
	loadServiceSetting(cfg.Sub("service"))
	loadSecret(cfg.Sub("security"))
	loadTimeFrom(cfg)

	return nil
}
