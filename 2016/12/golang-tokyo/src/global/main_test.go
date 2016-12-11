package main

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cases := []struct {
		path   string
		expect *Config
	}{
		{
			path: "test-fixture/golang-tokyo1.json",
			expect: &Config{
				Name: "deeeet",
			},
		},

		{
			path: "test-fixture/golang-tokyo2.json",
			expect: &Config{
				Name: "tcnksm",
			},
		},
	}

	for _, tc := range cases {
		config, _ := LoadConfig(tc.path)
		if !reflect.DeepEqual(config, tc.expect) {
			t.Fatalf("LoadConfig %#v, want %#v", config, tc.expect)
		}
	}
}
