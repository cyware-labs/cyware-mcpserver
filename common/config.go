package common

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

// Server holds configuration settings for the MCP server,
// including the mode (e.g., "stdio", "sse") and the port it listens on.
type Server struct {
	MCPMode string `mapstructure:"mcp_mode"`
	Port    string `mapstructure:"port"`
}

// Auth defines the authentication configuration for an application.
// It supports different auth types like "basic", "token".
type Auth struct {
	Type      string `mapstructure:"type"`
	Username  string `mapstructure:"username,omitempty"`
	Password  string `mapstructure:"password,omitempty"`
	Token     string `mapstructure:"token,omitempty"`
	AccessID  string `mapstructure:"access_id,omitempty"`
	SecretKey string `mapstructure:"secret_key,omitempty"`
}

// Application defines the configuration for an external application,
// including its base URL and authentication credentials.
type Application struct {
	BASE_URL string `mapstructure:"base_url"`
	Auth     Auth   `mapstructure:"auth"`
}

// Config holds the top-level configuration structure loaded from a YAML file.
// It includes all application definitions and the server settings.
type Config struct {
	Applications map[string]Application `mapstructure:"applications"`
	Server       Server                 `mapstructure:"server"`
}

// GetDomain extracts and returns the scheme and host (i.e., domain)
// from the given base URL. For example, "https://example.com/api" -> "https://example.com".
func GetDomain(baseURL string) string {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		return ""
	}

	// Construct the domain part (scheme + host)
	domain := parsedURL.Scheme + "://" + parsedURL.Host
	return domain
}

// Load reads and parses the YAML config file at the given path,
// and returns a populated Config struct.
// It returns an error if the file can't be read or parsed.
func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}
