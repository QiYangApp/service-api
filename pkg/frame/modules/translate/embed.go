//go:build go1.16
// +build go1.16

package translate

import (
	"embed"
)

type EmbedLoader struct {
	FS embed.FS
}

func (c *EmbedLoader) LoadMessage(path string) ([]byte, error) {
	return c.FS.ReadFile(path)
}
