package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	areaTests := []struct {
		name string
		s Shape
		want float64
	}{
		{name: "Rectangle", s: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", s: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", s: Triangle{12, 6}, want: 36.0},
	}

	// table driven tests
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.s, tt.want)
		})
	}

}