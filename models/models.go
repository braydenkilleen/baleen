package models

import (
	"database/sql"
	"log"
	"net/url"
)

// DB global database connection var. TODO: use DI instead of global
var DB *sql.DB

// Item represents a stored item.
type Item struct {
	ID      int
	URL     string
	Title   string
	Created string
	Updated string
	// Excerpt string
	// Content string
	// Tags     Tags
	// Archived bool
}

// AllItems returns all items in items table
func AllItems() ([]Item, error) {
	rows, err := DB.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item

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
		_, err = DB.Exec(
			"INSERT INTO items (url, title, created, updated) VALUES($1, $2, date('now'), date('now'))",
			u.String(), u.Host)
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}
	}

}
