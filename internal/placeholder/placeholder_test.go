package placeholder

import "testing"

func TestReadyPlaceholder(t *testing.T) {
	if Ready {
		t.Fatalf("unexpected placeholder value")
	}
}
