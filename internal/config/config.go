package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Log      LogConfig      `yaml:"log"`
}

type ServerConfig struct {
	Port           int      `yaml:"port"`
	Mode           string   `yaml:"mode"`
	AppName        string   `yaml:"app_name"`
	AllowedOrigins []string `yaml:"allowed_origins"`
	ReadTimeout    int      `yaml:"read_timeout"`
	WriteTimeout   int      `yaml:"write_timeout"`
	RequestTimeout int      `yaml:"request_timeout"`
}

type DatabaseConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
	TimeZone string `yaml:"time_zone"`

	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`

	AutoMigrate bool `yaml:"auto_migrate"`
}

func (d *DatabaseConfig) DSN() string {
	switch d.Driver {
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			d.Host, d.Username, d.Password, d.Database, d.Port, d.SSLMode, d.TimeZone)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			d.Username, d.Password, d.Host, d.Port, d.Database, strings.ReplaceAll(d.TimeZone, "/", "%2F"))
	default:
		return ""
	}
}

type RedisConfig struct {
	Enabled      bool   `yaml:"enabled"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
	Console    bool   `yaml:"console"`
}

var GlobalConfig *Config

func Init(configPath string) error {
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	GlobalConfig = &Config{}
	if err := yaml.Unmarshal(data, GlobalConfig); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	if GlobalConfig.Server.Port == 0 {
		GlobalConfig.Server.Port = 8080
	}
	if GlobalConfig.Server.Mode == "" {
		GlobalConfig.Server.Mode = "release"
	}
	if GlobalConfig.Server.RequestTimeout <= 0 {
		GlobalConfig.Server.RequestTimeout = 30
	}
	if len(GlobalConfig.Server.AllowedOrigins) == 0 {
		GlobalConfig.Server.AllowedOrigins = []string{"*"}
	}

	return nil
}

func GetConfig() *Config {
	return GlobalConfig
}