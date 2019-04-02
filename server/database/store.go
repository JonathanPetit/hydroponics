package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase(dbname string, host string, user string, password string, port int) *sql.DB {
	connStr := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		dbname, user, password, host, port,
	)
	fmt.Printf(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Printf("Unable to open connection to postgres:\n%v", err)
	}

	return db
}
