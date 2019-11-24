package internal_test

import (
	"testing"

	"github.com/NasSilverBullet/atc/internal"
)

func TestLanguagesConfig_GetLanguagePath(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		lc   internal.LanguagesConfig
		args args
		want string
	}{
		{"Python", internal.LanguagesConfig{"py": "python3", "go": "go run"}, args{"hoge.py"}, "python3"},
		{"Go", internal.LanguagesConfig{"py": "python3", "go": "go run"}, args{"hoge.go"}, "go run"},
		{"NoResult", internal.LanguagesConfig{"py": "python3", "go": "go run"}, args{"hoge.cpp"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lc.GetLanguagePath(tt.args.fileName); got != tt.want {
				t.Errorf("LanguagesConfig.GetLanguagePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
