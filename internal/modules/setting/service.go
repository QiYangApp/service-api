// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"github.com/spf13/viper"
)

var ServiceSetting = struct {
}{}

func loadServiceSetting(viper *viper.Viper) {
	//if err := viper.Unmarshal(ServiceSetting); err != nil {
	//	log.Client.Error("load service setting")
	//}
}
