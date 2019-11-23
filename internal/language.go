package internal

import (
	"path/filepath"
)

// GetLanguagePath return local laguange path
func GetLanguagePath(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".py":
		return "~/.pyenv/versions/pypy3.6-7.2.0/bin/pypy3"
	default:
		return ""
	}
}