package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"io"
	"flag"
)

// Inspired by: https://github.com/thbar/golang-playground/blob/master/download-files.go
func fetchURL(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, error := io.Copy(output, response.Body)
	if error != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
}

func main() {
	urlPtr := flag.String("url", "", "The base URL of the site to be downloaded. ex) http://google.com")

	flag.Parse()

	// If a URL was not specified, throw an error
	if *urlPtr == "" {
		fmt.Println("URL must not be empty. Provide a URL with the --url flag.")
	} else {
		fetchURL( *urlPtr )
	}
}