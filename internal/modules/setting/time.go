// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"frame/modules/log"
	"frame/util/timeutil"
	"time"

	"github.com/spf13/viper"
)

// DefaultUILocation is the location on the UI, so that we can display the time on UI.
var DefaultUILocation = time.Local

func loadTimeFrom(cfg *viper.Viper) {
	timeutil.DefaultUILocation = DefaultUILocation
	zone := cfg.GetString("time_zones")
	if zone != "" {
		var err error
		DefaultUILocation, err = time.LoadLocation(zone)
		if err != nil {
			log.Sugar().Fatalf("Load time zone failed: %v", err)
		} else {
			log.Sugar().Infof("Default time zones is %v", zone)
		}
	}
	if DefaultUILocation == nil {
		DefaultUILocation = time.Local
	}
}
