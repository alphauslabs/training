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
		Id:    "433178",
		Value: "Hi!, Lord Kelvin B. Abellana here! :)",
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("X-Request", "Custom Header Value")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(respBody))
}
