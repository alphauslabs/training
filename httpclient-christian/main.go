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

func sendRequest(url string, data msgT) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("X-Request", "SampleHeader")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func main() {
	url := "http://localhost:8080"
	data := msgT{
		Id:    "0",
		Value: "Hello World!",
	}

	response, err := sendRequest(url, data)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	fmt.Println(string(response))
}
