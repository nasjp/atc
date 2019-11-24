package internal

import (
	"path/filepath"
)

// GetLanguagePath return local laguange path
// TODO yamlの設定ファイルから読めるようにする
// TODO 実行にコンパイルの必要な言語にも対応する
func GetLanguagePath(fileName string) string {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".py":
		return "~/.pyenv/versions/pypy3.6-7.2.0/bin/pypy3"
	default:
		return ""
	}
}
