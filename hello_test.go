package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Say hello to a person", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world!' when the string is empty", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q does not equal want %q", got, want)
	}
}
