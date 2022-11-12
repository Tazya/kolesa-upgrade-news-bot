package config

import (
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port string
}

func ConfigPort() string {
	cfg := &Config{}
	_, err_toml := toml.DecodeFile("config/port.toml", cfg)

	_, err_conv := strconv.Atoi(cfg.Port)

	if err_toml != nil || err_conv != nil {
		return "8888"
	} else {
		return cfg.Port
	}
}
