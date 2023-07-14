package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type msgT struct {
	Id string "json:'id'"
	Value string "json:'value'"
}

func main() {
	url := "http://35.229.150.177:8080"

	msg := msgT {
		Id: "333",
		Value: "Hello, server! I am Jovi!",
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error in Marshaling json: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonMsg))
	if err != nil {
		log.Fatalf("Error in making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Connected successfully to the server!")
	} else {
		fmt.Println("Failed to connect to the server with status code: %v", resp.StatusCode)
	}
}
