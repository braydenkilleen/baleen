package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/braydenkilleen/baleen/models"
	"github.com/braydenkilleen/baleen/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup database
	var err error
	models.DB, err = sql.Open("sqlite3", "./tmp/test.db")
	if err != nil {
		log.Fatal(err)
	}

	// CLI
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		models.AddItems(addCmd.Args())
	case "list":
		listCmd.Parse(os.Args[2:])
		items, err := models.AllItems()
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range items {
			log.Printf("%s", i.URL)
		}
	case "serve":
		serveCmd.Parse(os.Args[2:])
		web.Serve()

	default:
		fmt.Println("expected `add` subcommand.")
		os.Exit(1)
	}

}
