package chanreq

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Job job struct
type Job struct {
	Payload Payload
}


//PayloadCollection payload
type PayloadCollection struct {
	Payloads []Payload `json:"data"`
}

//Payload struct
type Payload struct {
	RepoURL  string `json:"repoURL"`
	RepoName string `json:"repoName"`
}

//RequestParser parser
func (p *Payload) RequestParser() error {
	rsp, err := http.Get(p.RepoURL)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	defer rsp.Body.Close()
	responseData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	log.Println("Request builder: Scanning repo:\t", responseString)
	log.Print("\n\n\n\n\n")
	return nil
}
