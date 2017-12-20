package main

import (
	"fmt"
	"strings"
	"os"
	"net/http"
	"flag"
	"bytes"
	"io/ioutil"
	"path/filepath"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Inspired by: https://github.com/thbar/golang-playground/blob/master/download-files.go
func fetchURL( url string, rootURL string, pathToWrite string ) {
	tokens := strings.Split( url, "/" )
	fileName := tokens[len( tokens ) - 1]
	fmt.Println( "Downloading", url, "to", fileName )

	var pathBuffer bytes.Buffer

	pathBuffer.WriteString(pathToWrite)
	pathBuffer.WriteString("/")
	pathBuffer.WriteString(fileName)

	filePath := pathBuffer.String()

	// If this URL is the root URL, the final filename should be index.html

	output, err := os.Create( filePath )
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

    strToWrite := []byte(contentsOfPage)

	if err != nil {
		fmt.Println( "Error while downloading", url, "-", err )
		return
	}
	defer response.Body.Close()

    writeError := ioutil.WriteFile( filePath, strToWrite, 0755)
    check(writeError)

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
}

func main() {
	urlsToParse := make( []string, 0 )

	urlPtr := flag.String( "url", "", "(required) The base URL of the site to be downloaded. ex) http://google.com" )
	filePtr := flag.String( "path", "", "(required) The path where the downloaded site will be stored." )

	flag.Parse()

	// If a URL was not specified, throw an error
	if *urlPtr == "" {
		fmt.Println( "URL must not be empty. Provide a URL with the --url flag." )
		return
	}

	// If a file path was not specified, throw an error
	if *filePtr == "" {
		fmt.Println( "The path param must not be empty. Provide a path with the --path flag." )
		return
	}

	// Set up the root for the parse, based on the provided directory
	urlsToParse = append( urlsToParse, *urlPtr )

	// Create a folder based on the path passed in as a command line arg
	os.Mkdir( *filePtr, 0755)
	pathToWrite, _ := filepath.Abs(*filePtr)

	// If present, parse a new link from the queue every 5 seconds

	// If queue is empty for 30 seconds, end parsing

	fetchURL( urlsToParse[0], *urlPtr, pathToWrite )
}