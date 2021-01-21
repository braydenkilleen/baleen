package main

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./tmp/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]

	switch args[0] {
	case "add":
		addBookmark(args[1])
	default:
		println("unkown cmd")
	}
}

func addBookmark(rawurl string) {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec(
		"INSERT INTO item (link, created) VALUES($1, date('now'))",
		u.String())
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	log.Printf("%v", result)
}
