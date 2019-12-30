package main

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	r float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func Area(circle Circle) float64 {
	return 0
}
