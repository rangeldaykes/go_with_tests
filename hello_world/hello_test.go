package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	got := HelloWithName("world")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello_subTests(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWithName("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloWithName("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestChechLanguage(t *testing.T) {

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloWithNameAndLanguage("Elodie", "Spanish")
		want := "Hola, Elodie"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in French", func(t *testing.T) {
		got := HelloWithNameAndLanguage("Jaquin", "French")
		want := "Bonjour, Jaquin"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
