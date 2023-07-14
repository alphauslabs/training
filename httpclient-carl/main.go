package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type msgT struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func main() {
	url := "http://localhost:8080"

	
	msg := msgT{
		Id:    "101",
		Value: "Hello, server! Hope your are doing well!",
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}


	fmt.Println(string(body))
}
