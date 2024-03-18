// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"framework/log"
	"github.com/spf13/viper"
	"time"
)

// DefaultUILocation is the location on the UI, so that we can display the time on UI.
var DefaultUILocation = time.Local

func loadTimeFrom(cfg *viper.Viper) {
	zone := cfg.GetString("default_ui_location")
	if zone != "" {
		var err error
		DefaultUILocation, err = time.LoadLocation(zone)
		if err != nil {
			log.Client().Sugar().Fatal("Load time zone failed: %v", err)
		} else {
			log.Client().Sugar().Info("Default UI Location is %v", zone)
		}
	}
	if DefaultUILocation == nil {
		DefaultUILocation = time.Local
	}
}