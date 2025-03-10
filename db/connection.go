package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "ep-curly-leaf-a5gfc9g2-pooler.us-east-2.aws.neon.tech"
	port     = 5432
	user     = "restAPI_owner"
	password = "npg_5wzi0fQAFPIB"
	dbname   = "restAPI"
)

func DbConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
