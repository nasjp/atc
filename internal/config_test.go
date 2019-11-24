package internal_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/atc/internal"
)

const testConfigPath = "test/config.yml"

func setupTestNewConfig(t *testing.T) func(t *testing.T) {
	tmp := internal.ConfigPath

	internal.ConfigPath = testConfigPath

	config := `credentials:
  user_name: "user name"
  password: "password"
languages:
  py: "python3"
  go: "go run"`

	f, err := os.Create(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	f.WriteString(config)

	return func(t *testing.T) {
		p, err := filepath.Abs(testConfigPath)
		if err != nil {
			t.Fatal(err)
		}

		if err := os.Remove(p); err != nil {
			t.Fatal(err)
		}

		internal.ConfigPath = tmp
	}
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *internal.Config
	}{
		{"Success", &internal.Config{internal.CredentialsConfig{"user name", "password"}, internal.LanguagesConfig{"py": "python3", "go": "go run"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setupTestNewConfig(t)(t)
			got, err := internal.NewConfig()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
