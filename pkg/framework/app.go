package framework

import (
	"framework/cmd"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
)

// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

var App = cli.NewApp()

const Version = "0.0.1"
const Name = "QiYang"
const Usage = "QiYang"

func NewApp() {
	App.Name = Name
	App.Usage = Usage
	App.Version = Version
	App.Commands = []cli.Command{
		cmd.Web,
	}

	if err := App.Run(os.Args); err != nil {
		zap.S().Fatalf("Failed to start application: %v", err)
	}
}
