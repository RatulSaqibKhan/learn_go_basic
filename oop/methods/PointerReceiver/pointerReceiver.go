package main

import "fmt"

type Point struct {
	X int
	Y int
}
func (p *Point) MovePoint(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{12,10}
	p.MovePoint(2, 3)
	fmt.Printf("%+v\n", p)
}