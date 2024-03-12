package utils

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

var Path = new(PathHelpers).init()

type PathHelpers struct {
	RootPath       string
	StoragePath    string
	LogPath        string
	I18Path        string
	ControllerPath string
	ApiPath        string
	PublicPath     string
	ResourcePath   string
}

// https://www.pudongping.com/posts/7c9ed3a3.html
func (p *PathHelpers) init() *PathHelpers {
	p.RootPath = p.GetCurrentRunPath()
	p.ApiPath = p.JoinCurrentRunRootPath("/api")
	p.ControllerPath = p.Join(p.ApiPath, "/http/controller")
	p.StoragePath = p.JoinCurrentRunRootPath("/storage")
	p.ResourcePath = p.JoinCurrentRunRootPath("/resources")
	p.PublicPath = p.JoinCurrentRunRootPath("/public")
	p.LogPath = p.Join(p.StoragePath, "/log")
	p.I18Path = p.Join(p.ResourcePath, "/lang")

	return p
}

func (p PathHelpers) GetCurrentRunPath() string {
	path, err := os.Getwd()
	if err != nil || path == "" {
		zap.S().Fatalf("path get error %v", err.Error())
	}

	return path
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
