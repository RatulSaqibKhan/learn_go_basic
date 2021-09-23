package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	getContentType := resp.Header.Get("Content-Type")

	if getContentType == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}
	return getContentType, nil
}

func main() {
	getContentType, errors := contentType("https://www.codewithrsk.com/")

	if errors != nil {
		fmt.Printf("ERROR: %s\n", errors)
	} else {
		println(getContentType)
	}
}
