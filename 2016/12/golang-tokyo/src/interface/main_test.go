package main

import "testing"

// START OMIT
type TestDB struct {
	data string
}

func NewTestDB(data string) DB {
	return &TestDB{
		data: data,
	}
}

func (d *TestDB) Get(key string) (string, error) {
	return d.data, nil
}

// END OMIT

func TestProcess(t *testing.T) {
	want := "golang-tokyo"
	testDB := NewTestDB(want)
	got, err := process(testDB)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("process=%s, want = %s", got, want)
	}
}
