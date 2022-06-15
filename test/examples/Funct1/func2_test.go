package function

import "testing"

func TestBill(t *testing.T) {
	got := 90 * 6
	want := 540
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
