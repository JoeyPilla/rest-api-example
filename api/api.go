package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "joeypilla"
	password = ""
	dbname   = "restapiexample"
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	err = InitalizeDB(db)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully initalized articles table")
}

func InitalizeDB(db *sql.DB) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, `CREATE TABLE articles (
	id integer NOT NULL PRIMARY KEY,
	content TEXT,
	authorID integer NOT NULL,
	title varchar(40) NOT NULL,
	)`)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
