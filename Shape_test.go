package main

//Go Test file for unit test cases written for Shape.go 

import ("testing")

func TestAreaCircle(t *testing.T) {
	c := Circle{0, 0, 10}
	actualArea := computeArea(c)
	expectedArea := c.area()

	if expectedArea != actualArea {
		t.Errorf("Circle's area: expected %f, actual %f", expectedArea, actualArea)
	}
}

func TestAreaRectangle(t *testing.T) {
	r := Rectangle{10, 10, 15, 25}
	actualArea := computeArea(r)
	expectedArea := r.area()

	if expectedArea != actualArea {
		t.Errorf("Circle's area: expected %f, actual %f", expectedArea, actualArea)
	}
}

func TestPerimeterCircle(t *testing.T) {
	c := Circle{0, 0, 10}
	actualPerimeter := computePerimeter(c)
	expectedPerimeter := 62.83185307179586

	if expectedPerimeter != actualPerimeter {
		t.Errorf("Circle's Perimeter: expected %f, actual %f", expectedPerimeter, actualPerimeter)
	}
}

func TestPerimeterRectangle(t *testing.T) {
	r := Rectangle{10, 10, 15, 25}
	actualPerimeter := computePerimeter(r)
	expectedPerimeter := 40.0

	if expectedPerimeter != actualPerimeter {
		t.Errorf("Rectangle's Perimeter: expected %f, actual %f", expectedPerimeter, actualPerimeter)
	}
}
