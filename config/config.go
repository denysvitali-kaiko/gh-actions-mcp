package config

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Token     string `mapstructure:"token"`
	RepoOwner string `mapstructure:"repo_owner"`
	RepoName  string `mapstructure:"repo_name"`
	LogLevel  string `mapstructure:"log_level"`
}

var log = logrus.New()

func SetLogger(l *logrus.Logger) {
	log = l
}

func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("log_level", "info")
	v.SetDefault("token", "")

	// Environment variables
	v.SetEnvPrefix("GH")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME/.config/gh-actions-mcp")
		v.AddConfigPath("/etc/gh-actions-mcp")
	}

	// Try to read config file, ignore errors if not found
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Warnf("Error reading config file: %v", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Override with environment variable if set
	if token := v.GetString("token"); token != "" {
		cfg.Token = token
	}

	log.Debugf("Loaded config: owner=%s, repo=%s", cfg.RepoOwner, cfg.RepoName)
	return &cfg, nil
}

func (c *Config) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("GitHub token is required (set GITHUB_TOKEN env var or token in config)")
	}
	if c.RepoOwner == "" {
		return fmt.Errorf("repository owner is required (set repo_owner in config or use --repo)")
	}
	if c.RepoName == "" {
		return fmt.Errorf("repository name is required (set repo_name in config or use --repo)")
	}
	return nil
}
