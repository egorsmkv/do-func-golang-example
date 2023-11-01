package main

import (
	"testing"
)

// TestMain tests our function
func TestMain(t *testing.T) {
	in := Request{
		Name: "Yehor",
	}

	want := Response{
		Body: "Hello, Yehor",
	}

	msg, err := Main(in)
	if msg.Body != want.Body || err != nil {
		t.Fatalf(`Main("Yehor") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
