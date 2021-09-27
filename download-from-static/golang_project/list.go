package main

import (
	"io/ioutil"
	"net/http"
)

func html_dump (s string) (string, error) {
	
	start := "http://127.0.0.1:8080/"
	
	url := s
	if url == "" {
		url = start
	}
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		return "", err
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// show the HTML code as a string
	return string(html), nil

}