package config_base

import "github.com/BurntSushi/toml"

type BaseConfig struct {
	title string
	Base base	`toml:base`
}

type base struct {
	Port	int `toml:"port"`
	Host	string `toml:"host"`
}

func Config() BaseConfig {
	var config BaseConfig
	filePath := "config.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	return config
}