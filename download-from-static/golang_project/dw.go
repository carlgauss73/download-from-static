package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

func dw(dir string) error{

	fullURLFile = dir

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		return err
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	

	defer file.Close()
	return nil;
}