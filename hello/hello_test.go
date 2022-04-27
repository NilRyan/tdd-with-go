package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Philip")
	want := "Hello, Philip."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
