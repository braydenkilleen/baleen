package main

import (
	"database/sql"
	"log"

	"github.com/braydenkilleen/baleen/database"
	_ "github.com/mattn/go-sqlite3"
)

// App wrapper for the application.
type App struct {
	Name string
	Usage string
	DB *sql.DB
}

func main() {
	// Setup database
	// var err error
	// models.DB, err = sql.Open("sqlite3", "./tmp/test.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Initialize the database
	err := database.InitDB("tmp/test.db")
    if err != nil {
        log.Fatal(err)
	}

	items, err := database.AllItems()
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range items {
			log.Printf("%s", i.URL)
		}


	// CLI
	// addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	// serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	// switch os.Args[1] {
	// case "add":
	// 	addCmd.Parse(os.Args[2:])
	// 	// models.AddItems(addCmd.Args())
	// case "list":
	// 	listCmd.Parse(os.Args[2:])
		
	// case "serve":
	// 	serveCmd.Parse(os.Args[2:])
	// 	web.Serve()

	// default:
	// 	fmt.Printf("%q is not a valid command.\n", os.Args[1])
	// 	os.Exit(1)
	// }

}
