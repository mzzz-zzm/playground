package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		buf := bytes.Buffer{}
		Greet(&buf, "Chris")

		got := buf.String()
		want := "Hello, Chris"
		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
