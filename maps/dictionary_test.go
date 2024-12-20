package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinitions(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add a word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Add(word, definition)

		assertErrors(t, err, nil)
		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		got := dictionary.Add(word, definition)

		assertErrors(t, got, ErrAlreadyExists)
		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("update a word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		new_definition := "new definitionn"

		dictionary := Dictionary{word: definition}

		dictionary.Update(word, new_definition)

		assertDefinitions(t, dictionary, word, new_definition)

	})
}
