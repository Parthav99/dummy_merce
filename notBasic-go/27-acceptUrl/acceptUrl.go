// Extend (26) to accept URL as a command line argument instead of a hardcoded URL within
// the program.
package main

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	//Fetching a url
	fileUrl := os.Args[1]

	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments provided. Please provide only one argument.")
		os.Exit(1)
	}

	//Skipping ssl verification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	//Validating url
	parsedUrl, err := url.Parse(fileUrl)
	handleError(err)

	//Accessing url
	validUrl, err := client.Get(parsedUrl.String())
	handleError(err)
	defer validUrl.Body.Close()

	//Reading html content from url
	urlHtmlContent, err := ioutil.ReadAll(validUrl.Body)
	handleError(err)

	//Creating zip file
	newFile, err := os.Create("./merce-homepage.html" + ".zip")
	handleError(err)
	defer newFile.Close()

	//Creating zip writer
	gzipWrite := gzip.NewWriter(newFile)
	gzipWrite.Write(urlHtmlContent)
	fmt.Println("File successfully compressed")
	gzipWrite.Close()

	zipInfo, err := os.Stat(newFile.Name())
	handleError(err)

	//Printing size of files in bytes
	fmt.Println("Size of original file in bytes:", len(urlHtmlContent))
	fmt.Println("Size of compressed file in bytes:", zipInfo.Size())

}

//Handles errors
func handleError(err error) {
	if err != nil {
		log.Fatal("\n", err)
	}
}
