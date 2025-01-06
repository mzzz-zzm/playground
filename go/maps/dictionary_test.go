package maps

import (
	"testing"
)

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
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		err := d.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Update(word, "new definition")
		assertDefinition(t, dictionary, word, "new definition")
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{}

		err := d.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "test definition"}
		d.Delete(word)
		_, err := d.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		err := d.Delete(word)
		assertError(t, err, ErrNotFound)
	})	
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}