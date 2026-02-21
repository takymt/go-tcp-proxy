package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Args
	}{
		{
			name: "explicit",
			args: []string{"--listen", "127.0.0.1:9000", "--backends", "127.0.0.1:9001,127.0.0.1:9002"},
			want: Args{"127.0.0.1:9000", []string{"127.0.0.1:9001", "127.0.0.1:9002"}},
		},
		{
			name: "default",
			args: nil,
			want: Args{defaultListenAddr, strings.Split(defaultBackends, ",")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArgs(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestRoundRobin(t *testing.T) {
	r := RoundRobin{
		backends: []string{"127.0.0.1:9001", "127.0.0.1:9002", "127.0.0.1:9003"},
		current:  0,
	}
	want := "127.0.0.1:9002"

	got := r.Next()

	if got != want {
		t.Errorf("got=%v, want=%v", got, want)
	}
}
