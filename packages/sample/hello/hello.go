package main

import (
	"fmt"
	// "math/rand"
	// "time"

	log "github.com/sirupsen/logrus"
)

// const (
// 	// upperBound is the upper bound of the random int range
// 	upperBound = 101
// )

// func init() {
// 	rand.Seed(time.Now().Unix())
// }

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

func Main(in Request) (*Response, error) {
	// An example of using the logrus logger
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	return &Response{
		Body: fmt.Sprintf("Hello, %s", in.Name),
	}, nil
}
