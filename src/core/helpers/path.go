package helpers

import (
	"fmt"
	"os"
	"strings"
)

var PathInstance *Path = new(Path).init()

type Path struct {
	rootPath string
}

// https://www.pudongping.com/posts/7c9ed3a3.html
func (p *Path) init() *Path {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	p.rootPath = path

	return p
}

func (p *Path) setCurrentRunRootPath(path string) *Path {
	p.rootPath = path

	return p
}

func (p Path) GetCurrentRunRootPath() string {
	return p.rootPath
}

func (p Path) GetCurrentRunPath() string {
	if p.rootPath == "" {
		return ""
	}

	if strings.Contains(p.rootPath, "src") {
		return p.rootPath
	}

	return p.rootPath + "/src"
}

func (p Path) JoinCurrentRunPath(path string) string {
	return fmt.Sprintf("%s/%s", p.GetCurrentRunPath(), path)
}

func (p Path) JoinCurrentRunRootPath(path string) string {
	return fmt.Sprintf("%s/%s", p.GetCurrentRunRootPath(), path)
}
