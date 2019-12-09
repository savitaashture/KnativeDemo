package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func secHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "response from B microservice")
}

func knHandler(w http.ResponseWriter, r *http.Request) {
	var (
		res *http.Response
		err error
	)
	res, err = http.Get("http://third.test.svc.cluster.local/c")
	if err != nil {
		log.Fatal("Error from client get", err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error from ioutil", err)
	}
	fmt.Fprintf(w, string(resp))
}

func main() {
	log.Print("microservice B started")
	http.HandleFunc("/b", secHandler)
	http.HandleFunc("/bc", knHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
