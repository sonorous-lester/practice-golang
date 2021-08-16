package map_practice

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary {
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("new word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		dictionary.Add(key, val)
		assertDefinition(t, dictionary, key, val)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		val := "this is not a test"
		want := ErrExists
		got := dictionary.Add(key, val)
		assertError(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	key := "test"
	old := "this is just a test"
	update := "this is just a update"
	dic := Dictionary{}

	t.Run("existing word", func(t *testing.T) {
		dic.Add(key, old)
		err := dic.Update(key, update)
		assertError(t, err, nil)
		assertDefinition(t, dic, key, update)
	})

	t.Run("new word", func(t *testing.T) {
		key := "hello"
		update := "hi, there"
		err := dic.Update(key, update)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	val := "this is just a test"
	dic := Dictionary{}
	dic.Add(key, val)
	dic.Delete(key)
	_, err := dic.Search(key)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t *testing.T, d Dictionary, key, val string){
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	if got != val {
		t.Errorf("got '%s' want '%s'", got, val)
	}
}

func assertError(t *testing.T, got, want error){
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
