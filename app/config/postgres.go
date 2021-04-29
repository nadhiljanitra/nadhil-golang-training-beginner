package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitPostgres() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUserName := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")

	pgDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUserName, dbPassword, dbName)

	fmt.Printf(pgDsn)

	db, err := sql.Open("postgres", pgDsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully connected to db!\n")
}
