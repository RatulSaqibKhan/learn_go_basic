package main

import (
	"fmt"
	"log"
	"math"
)

// Point is a 2D point
type Point struct {
	X int
	Y int
}

// Move the point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

// Move the center of the square
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (square *Square) Area() int {
	squareLength := float64(square.Length)
	area := int(math.Pow(squareLength, 2))
	
	return area
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be positive")
	}
	square := &Square {
		Center: Point{x,y}, 
		Length: length,
	}

	return square, nil
}

func main() {
	s, err := NewSquare(1,2,12)

	if err != nil {
		log.Fatalf("Error can't create the square. %s\n", err)
	}

	s.Move(2,3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}