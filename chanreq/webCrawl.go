package chanreq

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//CrawlRequest web crawling
func CrawlRequest() {
	ch := make(chan string)
	for i := 0; i < 1000; i++ {
		go crawlRequest("https://github.com/Step1N/goWebApp", ch)
		fmt.Println(<-ch)
	}

}

func crawlRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}
