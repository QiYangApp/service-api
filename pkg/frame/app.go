package frame

import (
	"frame/cmd"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var App = cli.NewApp()

const Version = "0.0.1"
const Name = "Frame"
const Usage = "Frame"

func NewApp() {
	App.Name = Name
	App.Usage = Usage
	App.Version = Version

	App.Commands = []*cli.Command{
		cmd.WebCommand,
	}

	if err := App.Run(os.Args); err != nil {
		zap.S().Fatalf("Failed to start application: %v", err)
	}
}
