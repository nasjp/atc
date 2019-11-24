package internal_test

import (
	"io/ioutil"
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

func setupTestExistConfig(t *testing.T) func(t *testing.T) {
	tmp := internal.ConfigPath

	internal.ConfigPath = testConfigPath

	f, err := os.Create(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

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

func setupTestExistConfigFail(t *testing.T) func(t *testing.T) {
	tmp := internal.ConfigPath

	internal.ConfigPath = testConfigPath

	return func(t *testing.T) {
		internal.ConfigPath = tmp
	}
}

func TestExistConfig(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(t *testing.T) func(t *testing.T)
		want      bool
	}{
		{"True", setupTestExistConfig, true},
		{"False", setupTestExistConfigFail, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			defer tt.setupFunc(t)(t)
			got, err := internal.ExistConfig()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("ExistConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupTestGenerateConfig(t *testing.T) func(t *testing.T) {
	tmp := internal.ConfigPath

	internal.ConfigPath = testConfigPath

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

func TestGenerateConfig(t *testing.T) {
	defer setupTestGenerateConfig(t)(t)

	if err := internal.GenerateConfig(); err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	f, err := os.Open(testConfigPath)
	if err != nil {
		t.Errorf("config file was not created")
		return
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if want, got := internal.ConfigFileTemplate, string(b); want != got {
		t.Errorf("GenerateConfig() => config file text %v, want %v", got, want)
	}
}
