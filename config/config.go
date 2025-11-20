package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Version      string
	ServiceName  string
	Port         string
	JwtSecureKey string
	DBConfig     DBConfig
}

func loadConfig() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("VERSION not set in environment")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in environment")
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("PORT not set in environment")
		os.Exit(1)
	}

	if _, err := fmt.Sscanf(port, "%d", new(int)); err != nil {
		fmt.Println("PORT must be a valid number")
		os.Exit(1)
	}

	jwtSecurekey := os.Getenv("JWT_SECURE_KEY")

	if jwtSecurekey == "" {
		fmt.Println("JWT Secure Key not set in environment")
		os.Exit(1)
	}

	fmt.Printf("Configuration loaded: \nVersion=%s,\nServiceName=%s,\nPort=%s\n", version, serviceName, port)

	// db

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		fmt.Println("DB_HOST not set in environment")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		fmt.Println("DB_PORT not set in environment")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		fmt.Println("DB_USER not set in environment")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		fmt.Println("DB_PASSWORD not set in environment")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		fmt.Println("DB_NAME not set in environment")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbSSLMode == "" {
		fmt.Println("DB_SSLMODE not set in environment")
		os.Exit(1)
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		Port:         port,
		JwtSecureKey: jwtSecurekey,
	}

	configuration.DBConfig = DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  dbSSLMode,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
