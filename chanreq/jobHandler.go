package chanreq

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	//MaxLength length
	MaxLength = int64(1000)
	//MaxWorker number of worker
	MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue number of queue
	MaxQueue = os.Getenv("MAX_QUEUE")
)

//CrawlRequestHandler handler
func CrawlRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler: Starting crawling, %s, %s", MaxWorker, MaxQueue)
	numberOfWorker, err := strconv.Atoi(MaxWorker)
	if err != nil {
		//fmt.Errorf("number of worker is incorrect: %s", numberOfWorker)
		return
	}
	jobQueue := make(chan Job)
	dispatcher := NewDispatcher(numberOfWorker, jobQueue)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dispatcher.Run()

	var content = &PayloadCollection{}
	err = json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("Error in payload %v", err)
		return
	}

	gitURL := generateGitURL(content.Payloads)
	for _, gURL := range gitURL {
		job := Job{Payload: Payload{gURL, ""}}
		jobQueue <- job
	}

	w.WriteHeader(http.StatusOK)
}

func generateGitURL(payloads []Payload) []string {
	var parts []string
	host := "https://api.github.com/repos"
	for _, payload := range payloads {
		owner := strings.Split(payload.RepoURL, "/")[3]
		repoName := strings.Split(payload.RepoURL, "/")[4]
		var baseURL strings.Builder
		baseURL.WriteString(host)
		baseURL.WriteString("/" + owner)
		baseURL.WriteString("/" + repoName)
		baseURL.WriteString("/git/trees/master")
		parts = append(parts, baseURL.String())
	}

	return parts
}
