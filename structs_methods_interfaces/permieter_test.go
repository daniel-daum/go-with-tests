package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	length := 5.0
	width := 5.0

	got := Perimeter(length, width)
	want := 20.00

	if got != want {
		t.Errorf("Got '%.g' want '%.g' given '%.g' and '%.g'", got, want, length, width)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name: "rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, hasArea: 35.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
