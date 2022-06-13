package function

import "testing"

func TestDmart(t *testing.T) {
	got := DmartAdd(10, 12)
	want := 22
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
