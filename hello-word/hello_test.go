package main

import "testing"

func TestHello(t *testing.T)  {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("to a person", func (t *testing.T) {
		got := Hello("Facina", "")
		want := "Hello, Facina"
	
		assertCorrectMessage(t, got ,want)
	})

	t.Run("empty string", func (t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got ,want)
	})

	t.Run("in Brazil", func (t *testing.T) {
		got := Hello("Facina", "Brazil")
		want := "Olá, Facina"

		assertCorrectMessage(t, got ,want)
	})

	t.Run("in French", func (t *testing.T) {
		got := Hello("Facina", "French")
		want := "Bonjour, Facina"

		assertCorrectMessage(t, got ,want)
	})
}