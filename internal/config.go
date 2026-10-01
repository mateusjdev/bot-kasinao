package internal

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

type TelegramConfig struct {
	Token  string `mapstructure:"token"`
	ChatID int64  `mapstructure:"chatId"`
}

type ConfigOptions struct {
	MediaFolder    string         `mapstructure:"mediaFolder"`
	RuntimeFile    string         `mapstructure:"runtimeFile"`
	TelegramConfig TelegramConfig `mapstructure:"telegram"`
}

func InitConfig() (*ConfigOptions, error) {
	viperCfg := viper.New()
	viperCfg.SetConfigName("kasinao")
	viperCfg.SetConfigType("toml")
	if runtime.GOOS == "windows" {
		// TODO: Windows
		viperCfg.AddConfigPath("./")
	} else {
		viperCfg.AddConfigPath("/usr/local/etc/")
	}

	if err := viperCfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var cfg ConfigOptions
	if err := viperCfg.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct: %v", err)
	}

	return &cfg, nil
}
