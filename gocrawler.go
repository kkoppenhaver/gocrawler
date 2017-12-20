package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"io"
	"flag"
	"bytes"
)

// Inspired by: https://github.com/thbar/golang-playground/blob/master/download-files.go
func fetchURL( url string ) {
	tokens := strings.Split( url, "/" )
	fileName := tokens[len( tokens ) - 1]
	fmt.Println( "Downloading", url, "to", fileName )

	output, err := os.Create( fileName )
	if err != nil {
		fmt.Println( "Error while creating", fileName, "-", err )
		return
	}
	defer output.Close()

	response, err := http.Get( url )
	
	// Convert the HTML parsed from the link into a string
	buf := new( bytes.Buffer )
    buf.ReadFrom( response.Body )
    contentsOfPage := buf.String()

	if err != nil {
		fmt.Println( "Error while downloading", url, "-", err )
		return
	}
	defer response.Body.Close()

	_, error := io.Copy( output, response.Body )
	if error != nil {
		fmt.Println( "Error while downloading", url, "-", err )
		return
	}

	// Loop through each of the assets specified on the page and download them
	// to the correct path based on the root directory    

    // Loop through every link on the page and add them to the global link queue
	selectLinks( contentsOfPage )
}

func downloadAsset(assetPath string) {
	// Check if the folder where the asset belongs is already present on the filesystem

	// If not, create the folder

	// Download the asset into the available file path
}

func selectLinks(pageContents string) {
	// Parse the HTML to grab all links that match the root domain to recursively parse
	// by adding them to the global link queue
	fmt.Printf( pageContents )
}

func main() {
	urlsToParse := make( []string, 0 )

	urlPtr := flag.String( "url", "", "The base URL of the site to be downloaded. ex) http://google.com" )

	flag.Parse()

	// If a URL was not specified, throw an error
	if *urlPtr == "" {
		fmt.Println( "URL must not be empty. Provide a URL with the --url flag." )
		return
	}

	// Set up the root for the parse, based on the provided directory
	urlsToParse = append( urlsToParse, *urlPtr )

	// Create a folder named after the domain that was passed in

	// If present, parse a new link from the queue every 5 seconds

	// If queue is empty for 30 seconds, end parsing

	fetchURL(urlsToParse[0])
	
	fmt.Println( urlsToParse )
}