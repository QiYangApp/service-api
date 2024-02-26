package framework

import (
	"framework/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

const Version = "0.0.1"
const Name = "QiYang"
const Usage = "QiYang"

func NewApp() {
	app := cli.NewApp()
	app.Name = Name
	app.Usage = Usage
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.Web,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
