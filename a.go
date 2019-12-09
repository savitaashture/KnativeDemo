package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://second/b")
	if err != nil {
		log.Fatal("Error from client get", err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error from ioutil", err)
	}
	fmt.Fprintf(w, string(resp))
}

func knativeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://second/bc")
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
	log.Print("microservice A started")
	http.HandleFunc("/ab", envHandler)
	http.HandleFunc("/abc", knativeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
