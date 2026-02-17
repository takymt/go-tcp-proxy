package main

import "testing"

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"explicit", []string{"--listen", "127.0.0.1:9000"}, "127.0.0.1:9000"},
		{"default", []string{}, defaultListenAddr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}
