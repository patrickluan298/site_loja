package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	conexao := "user=postgres port=5432 dbname=site_loja password=753159 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
