package main

import "testing"

func TestParseArgs(t *testing.T) {
	expected := "127.0.0.1:9000"

	actual, err := parseArgs([]string{"--listen", expected})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestParseArgsDefault(t *testing.T) {
	actual, err := parseArgs([]string{})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual != defaultListenAddr {
		t.Errorf("actual=%v, defaultListenAddr=%v", actual, defaultListenAddr)
	}
}
