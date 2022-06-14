package function

import (
	"testing"
)

func TestDmart(t *testing.T) {
	got := DmartAdd(10, 12)
	want := 22
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestArea(t *testing.T) {

	t.Run("square", func(t *testing.T) {
		square := Square{12, 6}
		got := square.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
