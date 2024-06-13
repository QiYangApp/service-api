package path

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
)

var RunPath = sync.OnceValue(RunPathFun)()
var RunInternalPath = JoinPath(RunPath, "/internal")
var RunPublicPath = JoinPath(RunPath, "/public")
var RunResourcePath = JoinPath(RunPath, "/resources")
var RunStoragePath = JoinPath(RunPath, "/storage")
var RunTranslatePath = JoinPath(RunResourcePath, "/translate")

func RunPathFun() string {
	path, err := os.Getwd()
	if err != nil || path == "" {
		zap.S().Fatalf("path get error %v", err.Error())
		return ""
	}

	path, err = filepath.Abs(path)
	if err != nil {
		zap.S().Fatalf("path get error %v", err.Error())
		return ""
	}

	return path
}

func JoinPath(paths ...string) string {
	path := filepath.Join(paths...)
	path, _ = filepath.Abs(path)

	return path
}
