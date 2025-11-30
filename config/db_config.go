package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDBConfig() {

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

	configuration.DBConfig = DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  dbSSLMode,
	}
}
