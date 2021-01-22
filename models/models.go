package models

// Tag ...
type Tag struct {
	ID int
	Name string
}

// Item represents a stored item.
type Item struct {
	ID      int
	URL     string
	Title   string
	Created string
	Updated string
}

// // AddItems adds items to items table
// func AddItems(rawurls []string) {
// 	for _, rawurl := range rawurls {
// 		u, err := url.Parse(rawurl)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		_, err = DB.Exec(
// 			"INSERT INTO items (url, title, created, updated) VALUES($1, $2, date('now'), date('now'))",
// 			u.String(), u.Host)
// 		if err != nil {
// 			log.Fatalf("Error: %v", err)
// 			return
// 		}
// 	}
// }
