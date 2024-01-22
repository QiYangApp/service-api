package helpers

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

var Path = new(PathHelpers).init()

type PathHelpers struct {
	RootPath     string
	StoragePath  string
	LogPath      string
	ResourcePath string
}

// https://www.pudongping.com/posts/7c9ed3a3.html
func (p *PathHelpers) init() *PathHelpers {
	p.RootPath = p.GetCurrentRunPath()
	p.StoragePath = p.JoinCurrentRunRootPath("/storage")
	p.LogPath = p.JoinCurrentRunRootPath("/storage/log")
	p.ResourcePath = p.JoinCurrentRunRootPath("/resource")

	return p
}

func (p PathHelpers) GetCurrentRunPath() string {
	path, err := os.Getwd()
	if err != nil || path == "" {
		zap.S().Fatalf("path get error %v", err.Error())
	}

	if strings.Contains(path, "src") {
		return path
	}

	return path + "/src"
}

func (p PathHelpers) JoinCurrentRunRootPath(path string) string {
	return p.Join(p.RootPath, path)
}

func (p PathHelpers) Join(dir, path string) string {
	path, _ = filepath.Abs(dir + "/" + path)

	return path
}

func (p PathHelpers) JoinStroragePath(path string) string {
	path, _ = filepath.Abs(p.StoragePath + "/" + path)

	return path
}
