package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func addBookmark(rawurl string) {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Host)
}

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "add":
		addBookmark(args[1])
	default:
		println("unkown cmd")
	}
}
