// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/log"
	"github.com/spf13/viper"
)

var ServiceSetting = struct {
	RegisterConfirm bool `mapstructure:"register_confirm"`
}{}

func loadServiceSetting(viper *viper.Viper) {
	if err := viper.Unmarshal(ServiceSetting); err != nil {
		log.Client.Error("load service setting")
	}
}
