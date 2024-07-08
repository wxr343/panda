package global

import (
	"github.com/spf13/viper"
	logconfig "github.com/wxr343/logger/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"panda/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
var LogCfg logconfig.Configuration
