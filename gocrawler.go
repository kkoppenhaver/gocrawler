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

func downloadAsset(assetPath string) {
	// Check if the folder where the asset belongs is already present on the filesystem

	// If not, create the folder

	// Download the asset into the available file path
}

func selectLinks(url string) {
	// Parse the HTML to grab all links that match the root domain to recursively parse
}

func main() {
	urlPtr := flag.String("url", "", "The base URL of the site to be downloaded. ex) http://google.com")

	flag.Parse()

	// If a URL was not specified, throw an error
	if *urlPtr == "" {
		fmt.Println("URL must not be empty. Provide a URL with the --url flag.")
	} else {
		// Create a folder named after the domain that was passed in

		// Fetch URL to the root of the newly created folder
		fetchURL( *urlPtr )
	}

	// Parse any image tags on the page and pull those down into the proper directories 
}