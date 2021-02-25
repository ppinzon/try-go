package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	resp, err := http.Get("http://localhost:1323/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	if sb != "Hello, World!" {
		t.Errorf("Not Hello World")
	}
}
