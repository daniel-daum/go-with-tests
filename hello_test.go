package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello to a person", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("Got %q does not equal want %q", got, want)
		}
	})

	t.Run("Say 'Hello, world!' when the string is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		if got != want {
			t.Errorf("Got %q does not equal want %q", got, want)
		}
	})

}