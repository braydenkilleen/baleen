package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/braydenkilleen/baleen/database"
	"github.com/braydenkilleen/baleen/web"
	_ "github.com/mattn/go-sqlite3"
)

// App wrapper for the application.
type App struct {
	Name string
	Usage string
	DB *sql.DB
}

func main() {
	// Initialize the database
	err := database.InitDB("tmp/test.db")
    if err != nil {
        log.Fatal(err)
	}

	// ...
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		database.AddItems(addCmd.Args())
	case "list":
		listCmd.Parsed()
		items, err := database.AllItems()
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
		fmt.Printf("%q is not a valid command.\n", os.Args[1])
		os.Exit(1)
	}

}
