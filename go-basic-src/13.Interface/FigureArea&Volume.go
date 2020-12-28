package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	volume() float64
}

type cylinder struct {
	r, h float64
}

type cuboid struct {
	side1, side2, side3 float64
}

func (cyl cylinder) area() float64 {
	a := math.Pi*cyl.r*cyl.r*2 + 2*math.Pi*cyl.r*cyl.h
	return a
}

func (cyl cylinder) volume() float64 {
	v := math.Pi * cyl.r * cyl.r * cyl.h
	return v
}

func (cub cuboid) area() float64 {
	a := 2 * (cub.side1*cub.side2 + cub.side1*cub.side3 + cub.side2*cub.side3)
	return a
}

func (cub cuboid) volume() float64 {
	v := cub.side1 * cub.side2 * cub.side3
	return v
}

func main() {
	cy1 := cylinder{10, 10}
	cy2 := cylinder{4.2, 15.6}
	cu1 := cuboid{10.5, 20.2, 20}
	cu2 := cuboid{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(m ...geometry) {
	for _, fig := range m {
		fmt.Printf("%.2f, ", fig.area())
		fmt.Printf("%.2f\n", fig.volume())
	}
}
