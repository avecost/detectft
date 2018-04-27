package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	t := os.Args[1]

	// Open file
	f, err := os.Open(t)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content Type: " + contentType)

}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DetectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil

}
