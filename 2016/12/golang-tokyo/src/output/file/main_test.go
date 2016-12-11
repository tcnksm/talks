package main

import (
	"bytes"
	"testing"
)

func TestSaveConfig(t *testing.T) {

	cases := []struct {
		Config *Config
		Output string
	}{
		{&Config{Name: "deeeet"}, `{"name":"deeeet"}
`},
	}

	for _, tc := range cases {
		var buf bytes.Buffer // 🙆
		if err := Save(tc.Config, &buf); err != nil {
			t.Fatal(err)
		}

		if got, want := buf.String(), tc.Output; got != want {
			t.Fatalf("Output: %q, want %q", got, want)
		}
	}
}
