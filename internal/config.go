package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var ConfigPath = "~/.atc.yml"

type Config struct {
	CredentialsConfig `yaml:"credentials"`
	LanguagesConfig   `yaml:"languages"`
}

type LanguagesConfig map[string]string

type CredentialsConfig struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	p, err := filepath.Abs(ConfigPath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("can't find ~/.atc.yml. please make this file and write config.")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	c := &Config{}

	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, fmt.Errorf("can't parse ~/.atc.yml. please confirm yaml format.")
	}

	return c, nil
}
