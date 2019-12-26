package main

import "testing"

func TestHello(t *testing.T) {

	assertCurrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("hans", "")
		want := "Hello, hans!"
		assertCurrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCurrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCurrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("龙哥", "Chinese")
		want := "你好, 龙哥!"
		assertCurrectMessage(t, got, want)
	})
}
