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
	Id string "json:'id'"
	Value string "json:'value'"
}

func main() {
	url := "http://10.140.0.9:8080"

	msg := msgT{
		Id: "333",
		Value: "Hi, I am Jovi!",
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error in marshaling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Error in creating request: %v", err)
	}

	req.Header.Set("X-Request", "Send a greeting")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request: %v", err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error in reading response body: %v", err)
	}

	fmt.Println(string(respBody))
	fmt.Println("Connected to server!")
}
