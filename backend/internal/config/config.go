package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port		string // Server port
	DBHost		string
	DBPort		string
	DBUser		string
	DBPassword	string
	DBName		string
	DBSSLMode	string
	Env			string
}

// LoadConfig loads configuration from environment variables or a .env file.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using environment variables")
		return &Config{}, err
	}
	// Load .env file if it exists
	cfg := &Config {
		Port:      	getEnv("PORT", "8080"),
		DBHost:	 	getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		Env:        getEnv("ENV", "development"),
	}

	// Validate required fields
	requiredFields := map[string]string{
		"DBUser":		cfg.DBUser,
		"DBPassword":	cfg.DBPassword,
		"DBName":		cfg.DBName,
	}

	for field, value := range requiredFields {
		if value == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", field)
		}
	}
	return cfg, nil
}

// getEnv retrieves the value of the environment variable named by the key.
// It returns the value, or the defaultValue if the variable is not present.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// DatabaseURL constructs the database connection URL.
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
		)
}
