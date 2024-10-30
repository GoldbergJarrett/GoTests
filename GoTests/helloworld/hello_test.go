package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Jarrett")
	want := "Hello, Jarrett!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
