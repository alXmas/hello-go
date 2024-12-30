package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Xmas")
	want := "Hello, Xmas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
