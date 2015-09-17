package main

import ("math")

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x_coordinate float64
	y_coordinate float64
	radius float64
}

//This is a function
func circleArea(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

//This is a method
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}


type Rectangle struct {
	x1_coordinate, y1_coordinate, x2_coordinate, y2_coordinate float64
}

func (r Rectangle) area() float64 {
	length := distance(r.x1_coordinate, r.y1_coordinate, r.x1_coordinate, r.y2_coordinate)
	width  := distance(r.x1_coordinate, r.y1_coordinate, r.x2_coordinate, r.y1_coordinate)
	return length * width
}

func (r Rectangle) perimeter() float64 {
	length := distance(r.x1_coordinate, r.y1_coordinate, r.x1_coordinate, r.y2_coordinate)
	width  := distance(r.x1_coordinate, r.y1_coordinate, r.x2_coordinate, r.y1_coordinate)
	return 2 * (length + width)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func computeArea(s Shape) float64 {
	var area float64
    area = s.area()
    return area
}

func computePerimeter(s Shape) float64 {
	var perimeter float64
    perimeter = s.perimeter()
    return perimeter
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

