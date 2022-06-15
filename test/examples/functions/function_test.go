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

// making sure the test output is helpfull
func TestArea1(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shapes
		hasArea float64
	}{

		{name: "Triangle", shape: Triangle{Height: 14, Base: 7}, hasArea: 49.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
