package internal

import (
	"path/filepath"
)

// GetLanguagePath return local laguange path
// TODO 実行にコンパイルの必要な言語にも対応する
func (lc LanguagesConfig) GetLanguagePath(fileName string) string {
	ext := filepath.Ext(fileName)
	if ext == "" {
		return ""
	}

	return lc[ext[1:]]
}
