package helpers

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
)

var once sync.Once
var PathInstance *Path

type Path struct {
	rootPath string
}

// https://www.pudongping.com/posts/7c9ed3a3.html
func (p *Path) init() *Path {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	zap.S().Info("project root path ", path)

	p.rootPath = path

	return p
}

func (p *Path) setCurrentRunRootPath(path string) *Path {
	p.rootPath = path

	return p
}

func (p Path) getCurrentRunRootPath() string {
	return p.rootPath
}

func (p Path) getCurrentRunPath() string {
	if p.rootPath == "" {
		return ""
	}

	if strings.Contains(p.rootPath, "src") {
		return p.rootPath
	}

	return p.rootPath + "/src"
}

func (p Path) JoinCurrentRunPath(path string) string {
	return fmt.Sprintf("%s/%s", p.getCurrentRunPath(), path)
}
