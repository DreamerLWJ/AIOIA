package config

import "github.com/BurntSushi/toml"

type Config struct {
}

func Init(path string) (*Config, error) {
	cfg := &Config{}
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
