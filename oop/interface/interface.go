package main

import (
	"fmt"
	"log"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

func NewSquare(length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be positive.")
	}

	return &Square{Length: length}, nil
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2) 
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("radius must be positive")
	}

	return &Circle{Radius: radius}, nil
}

func sumAreas(shapes []Shape) float64 {
	totalArea := 0.0

	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	return totalArea
}

type Shape interface {
	Area() float64
}

func main() {
	s, sq_err := NewSquare(12)

	if sq_err != nil {
		log.Fatalf("Can't create the Square. %s\n", sq_err)
	}

	fmt.Printf("Area of the square: %.2f\n", s.Area())
	
	c, c_err := NewCircle(4)

	if c_err != nil {
		log.Fatalf("Can't create the Circle. %s.\n", c_err)
	}

	fmt.Printf("Area of the circle: %.2f\n", c.Area())

	shapes := []Shape{s,c}
	shapesTotalArea := sumAreas(shapes)
	fmt.Printf("Total summation of the area: %.2f\n", shapesTotalArea)
}