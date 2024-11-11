package main

import "testing"

func TestHell0(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q does not equal want %q", got, want)
	}
}