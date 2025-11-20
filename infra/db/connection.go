package db

import (
	"ecom_project/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString(config *config.DBConfig) string {

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)

	return connectionString
}

func NewDBConnection(config config.DBConfig) (*sqlx.DB, error) {

	dbSource := GetDBConnectionString(&config)

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
