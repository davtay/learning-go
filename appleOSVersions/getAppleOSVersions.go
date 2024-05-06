package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	body := getURLBody()
	macOS := readMacOS(body)
	iOS := readiOSiPadOS(body)
	fmt.Printf("Latest macOS version: %v\n", macOS)
	fmt.Printf("Latest iOS and iPadOS version: %v\n", iOS)
}

func getURLBody() string {
	res, err := http.Get("https://support.apple.com/en-us/HT201222")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return string(body[:])
}

func readMacOS(body string) string {
	macOSExpression := regexp.MustCompile(`(?:macOS\sis&nbsp;)(\d+(.\d+)+)`)
	macOS := macOSExpression.FindStringSubmatch(body)
	if macOS[1] != "" {
		return macOS[1]
	} else {
		return "Version not found"
	}
}

func readiOSiPadOS(body string) string {
	iOSExpression := regexp.MustCompile(`(?:iOS\sand\siPadOS\sis\s)(\d+(.\d+)+)`)
	iOS := iOSExpression.FindStringSubmatch(body)
	if iOS[1] != "" {
		return iOS[1]
	} else {
		return "Version not found"
	}
}
