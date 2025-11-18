package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version      string
	ServiceName  string
	Port         string
	JwtSecureKey string
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

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		Port:         port,
		JwtSecureKey: jwtSecurekey,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
