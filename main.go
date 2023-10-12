package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	pingPeriod = 1 * time.Second // time between pings
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./http_ping <URL>")
		return
	}
	url := os.Args[1]

	// Create a transport that skips SSL verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second, // sets a 5-second timeout for the request
	}

	var startTime time.Time
	firstRequest := true

	for {
		if firstRequest {
			startTime = time.Now()
			firstRequest = false
		}

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			cumulativeTime := time.Since(startTime)
			fmt.Printf("%s HTTP %d - %s after %v\n", url, resp.StatusCode, http.StatusText(resp.StatusCode), cumulativeTime)
			resp.Body.Close()
		}

		time.Sleep(pingPeriod)
	}
}
