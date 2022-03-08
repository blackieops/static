package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Webroot        string           `yaml:"webroot"`
	Headers        []ResponseHeader `yaml:"headers"`
	TrustedProxies []string         `yaml:"trusted_proxies"`
}

type ResponseHeader struct {
	Name string `yaml:"name"`
	Value string `yaml:"value"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
