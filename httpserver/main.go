package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type msgT struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var reply strings.Builder
	if r.Method != http.MethodPost { // request method should be POST
		fmt.Fprintf(&reply, "I need a POST method.\n")
		fmt.Fprintf(w, reply.String())
		return
	}

	xhdr := r.Header.Get("X-Request")
	if xhdr == "" {
		fmt.Fprintf(&reply, "I need an X-Request header. Can you provide one?\n")
		fmt.Fprintf(w, reply.String())
		return

	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	defer r.Body.Close()
	var msg msgT
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Fprintf(&reply, "Your request body has the wrong format.\n")
		fmt.Fprintf(&reply, "Try using the 'msgT' format.\n")
		fmt.Fprintf(&reply, "Error: %v", err)
		fmt.Fprintf(w, reply.String())
		return
	}

	fmt.Fprintf(&reply, "Your method is %v.\n", r.Method)
	fmt.Fprintf(&reply, "Your X-Request header is %v.\n", xhdr)
	fmt.Fprintf(&reply, "Your request body is %v.\n", string(body))
	fmt.Fprintf(&reply, "All good!")
	fmt.Fprintf(w, reply.String())
	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
