package main

import (
	"net/http"
	"testing"
)

// TestMain tests our function
func TestMain(t *testing.T) {
	in := Request{
		Name: "Yehor",
	}

	want := Response{
		Body:       "Hello, Yehor",
		StatusCode: http.StatusOK,
	}

	msg, err := Main(in)
	if err != nil {
		t.Fatalf(`Main("Yehor"), error %v`, err)
	}

	if msg.Body != want.Body {
		t.Fatalf(`Main("Yehor") = %s, wrong body, must be %s`, msg.Body, want.Body)
	}

	if msg.StatusCode != want.StatusCode {
		t.Fatalf(`Main("Yehor"), status code is %d, wrong status code, must be %d`, msg.StatusCode, want.StatusCode)
	}
}
