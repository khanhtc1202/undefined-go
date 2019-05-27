package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	endpointList = [2]string{"http://localhost:9453/api/apple", "http://localhost:9453/api/banana"}
)

func main() {
	for {
		seed := rand.Intn(5)
		requestTo(seed)
		time.Sleep(time.Duration(seed) * time.Second)
	}
}

func requestTo(seed int) {
	url := endpointList[seed%2]
	log.Printf("Request to %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}
