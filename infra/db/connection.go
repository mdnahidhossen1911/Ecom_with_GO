package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString() string {
	return "user=postgres password=2003 host=localhost port=5432 dbname=ecom sslmode=disable"
}

func NewDBConnection() (*sqlx.DB, error) {

	dbSource := GetDBConnectionString()

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
