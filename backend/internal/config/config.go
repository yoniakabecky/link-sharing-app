package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Upload   UploadConfig
}

// UploadConfig holds avatar upload configuration
type UploadConfig struct {
	Dir      string // directory for uploads (e.g. "uploads"); avatars saved under Dir/avatars/
	MaxBytes int64  // max file size in bytes (e.g. 2 * 1024 * 1024 = 2MB)
	BaseURL  string // base URL for building avatar_url in responses (e.g. "http://localhost:8080")
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Address string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	JWTKey   string
	JWTExp   int64
}

type JWTConfig struct {
	Key        string
	Exp        int64 // access token expiry in seconds (e.g. 900 = 15 min)
	RefreshExp int64 // refresh token expiry in seconds (e.g. 7 days)
}

// returns the Data Source Name for database connection
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		d.User, d.Password, d.Host, d.Port, d.Database)
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", ":8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "password"),
			Database: getEnv("DB_NAME", "link-sharing-app"),
		},
		JWT: JWTConfig{
			Key:        getEnv("JWT_KEY", "secret"),
			Exp:        getEnvAsInt("JWT_EXP", 60*15),              // 15 min default
			RefreshExp: getEnvAsInt("JWT_REFRESH_EXP", 60*60*24*7), // 7 days default
		},
		Upload: UploadConfig{
			Dir:      getEnv("UPLOAD_DIR", "uploads"),
			MaxBytes: getEnvAsInt("UPLOAD_MAX_BYTES", 2*1024*1024), // 2MB default
			BaseURL:  getEnv("API_BASE_URL", "http://localhost:8080"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}
