package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	c "goWebCrawler/chanreq"

	"github.com/urfave/negroni"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/crawl", c.CrawlRequestHandler)

	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
	start := time.Now()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
