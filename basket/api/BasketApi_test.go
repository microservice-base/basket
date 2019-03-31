package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	given := productAPIMethod()
	expect := "Hello, '/productapi'"

	if given != expect {
		t.Errorf("given '%s' ", given)
		t.Errorf("expect '%s' ", expect)
	}
}
