package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Request is the function's request struct
type Request struct {
	Name string `json:"name,omitempty"`
}

// Response is the function's response struct
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

// Main is the function's entrypoint
func Main(in Request) (*Response, error) {
	// An example of using the logrus logger
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// An example of returning a response
	return &Response{
		Body:       fmt.Sprintf("Hello, %s", in.Name),
		StatusCode: http.StatusOK,
	}, nil
}
