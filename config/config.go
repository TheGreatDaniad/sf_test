package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all application configurations.
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Email    EmailConfig    `mapstructure:"email"`
	Metrics  MetricsConfig  `mapstructure:"metrics"`
}

// AppConfig holds general app-related configurations.
type AppConfig struct {
	Port         int    `mapstructure:"port"`
	Version      string `mapstructure:"version"`
	DoMigrations bool   `mapstructure:"do_migrations"`
}

// DatabaseConfig holds database connection details.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

// EmailConfig holds email client configurations.
type EmailConfig struct {
	SMTPHost    string `mapstructure:"smtp_host"`
	SMTPPort    int    `mapstructure:"smtp_port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	SenderEmail string `mapstructure:"sender_email"`
}

// MetricsConfig holds Prometheus metrics configurations.
type MetricsConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Path    string `mapstructure:"path"`
	Port    int    `mapstructure:"port"`
}

// LoadConfig initializes the application configuration from file and environment variables.
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// Set default configurations
	v.SetDefault("app.port", 8080)
	v.SetDefault("metrics.enabled", true)
	v.SetDefault("metrics.path", "/metrics")
	v.SetDefault("metrics.port", 9090)

	// Automatically read environment variables (app-specific prefix)
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read configuration from file
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	// Unmarshal configurations into the struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
