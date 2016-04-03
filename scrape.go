package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// initialize the database
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Post VALUES (NULL, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// read the response
	resp, err := http.Get("http://www.vertabelo.com/blog")
	if err != nil {
		log.Fatal(err)
	}
	// we are always expected to close out the body of a response when using the net/http package.
	defer resp.Body.Close()

	// initiate GoQuery from the response
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}


	// find DOM containers with the blog_article CSS class and loop over selection
	doc.Find(".blog_article").Each(func(i int, s *goquery.Selection) {

		// extract individual fields with GoQuery helpers
		title := s.Find(".ba_header").Text()
		link, _ := s.Find(".ba_header a").Attr("href")
		author := s.Find(".ba_author_name").Text()
		date := s.Find(".ba_publication_date").Text()
		summary := s.Find(".ba_summary P").Text()

		// insert the post into the database
		_, err := stmt.Exec(author, title, date, link, summary)
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Completed import")

}