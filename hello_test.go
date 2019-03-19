package main

import "testing"

func TestHello(t *testing.T) {
    given := Hello()
    expect := "Hello World First Go App"

    if given != expect  {
	t.Errorf("given '%s' ", given )
	t.Errorf("expect '%s' ", expect )
    }
}
