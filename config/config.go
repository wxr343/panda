package config

import (
	logger "github.com/wxr343/logger/config"
)

type Configuration struct {
	App      App        `mapstructure:"app" json:"app" yaml:"app"`
	Log      logger.Log `mapstructure:"log" json:"log" yaml:"log"`
	Database Database   `mapstructure:"database" json:"database" yaml:"database"`
}
