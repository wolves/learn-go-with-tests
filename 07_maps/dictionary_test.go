package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"

    err := dict.Add(word, definition)

		assertError(t, err, nil)
    assertDefinition(t, dict, word, definition)
	})

  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dict := Dictionary{word: definition}

    err := dict.Add(word, "new test")

    assertError(t, err, ErrWordExists)
    assertDefinition(t, dict, word, definition)
  })

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dict := Dictionary{word: definition}
    newDef := "new definition"

    err := dict.Update(word, newDef)

    assertError(t, err, nil)
    assertDefinition(t, dict, word, newDef)
  })

	t.Run("new word", func(t *testing.T) {
    word := "test"
    definition := "this is just a test"
    dict := Dictionary{}

    err := dict.Update(word, definition)

    assertError(t, err, ErrWordDoesNotExist)
  })
}

func TestDelete(t *testing.T) {
  word := "test"
  dict := Dictionary{word: "test definition"}

  dict.Delete(word)

  _, err := dict.Search(word)
  if err != ErrNotFound {
    t.Errorf("Expected %q to be deleted", word)
  }
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
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
