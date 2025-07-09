package main

import (
	"net/http"
	"time"
)

// custom client
type Client struct {
	httpClient http.Client
}

// creating new client constructor which takes the time duration
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
