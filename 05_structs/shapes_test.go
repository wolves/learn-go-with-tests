package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
    name string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}

	//  checkArea := func(t testing.TB, shape Shape, want float64) {
	//    t.Helper()
	//    got := shape.Area()
	//
	//		if got != want {
	//			t.Errorf("got %.2f want %.2f", got, want)
	//		}
	//  }
	//
	//	t.Run("rectangles", func(t *testing.T) {
	//		rect := Rectangle{12, 6}
	//		checkArea(t, rect, 72.0)
	//	})
	//
	//  t.Run("circles", func(t *testing.T) {
	//    circle := Circle{10}
	//    checkArea(t, circle, 314.1592653589793)
	//  })
}
