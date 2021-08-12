package struct_method

import "testing"

func checkResult(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	checkResult(t, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct{
		name string
		shape Shape
		want float64
	}{
		{"Rectangle", Rectangle{12, 6	}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %.2f want %.2f", got, test.want)
			}
		})
	}
}