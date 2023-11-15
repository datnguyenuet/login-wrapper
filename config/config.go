package config

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"

	"login-wrapper/pkg/logging"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	HTTPPort    string `mapstructure:"HTTP_PORT"`
}

func LoadFromEnv(ctx context.Context) (*Config, error) {
	log := logging.FromContext(ctx)
	log.Infof("starting load config from ENV ...")

	var cfg *Config

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if _, err := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed()); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
