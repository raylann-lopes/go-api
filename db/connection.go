package db

import (
	"fmt"
	"database/sql"
)

const (
	host = "restAPI_owner"
	port = 5432
	user = "restAPI"
	password = "npg_5wzi0fQAFPIB"
	dbname = "restAPI"
)

func DbConnection () (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=&s" + 
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
