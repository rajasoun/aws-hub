package main

import (
	"log"
	"net/http"

	"github.com/rajasoun/aws-hub/test/examples/httpmock/https"
)

var client https.Client

func init() {
	client = &http.Client{}
}

func main() {
	Ping()
}

//this func will check the health of given url
func Ping() {
	url := "https://google.com"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Authorization", "test")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("website is not healthy")
	}
	log.Println("All good.")
}
