package database

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/braydenkilleen/baleen/models"
)

var db *sql.DB

// InitDB creates & opens a connection to db.
func InitDB(dsn string) (err error) {
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}

// AllItems returns all items in items table
func AllItems() ([]models.Item, error) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.Title, &item.URL, &item.Created, &item.Updated)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// AddItems adds items to items table
func AddItems(rawurls []string) {
	for _, rawurl := range rawurls {
		u, err := url.Parse(rawurl)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(
			"INSERT INTO items (url, title, created, updated) VALUES($1, $2, date('now'), date('now'))",
			u.String(), u.Host)
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}
	}
}