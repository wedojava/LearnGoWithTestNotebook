package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		want    float64
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
		}
	}
	// https://learnku.com/docs/learn-go-with-tests/structure-method-and-interface/6084#129bda
}
