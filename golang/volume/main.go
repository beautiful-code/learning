package main

import (
	"fmt"
	"math"
)

type Structure interface {
	volume() float64
}
type Cube struct {
	side float64
}
type Box struct {
	length, width, height float64
}
type Sphere struct {
	radius float64
}

func (sphere Sphere) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(sphere.radius, 3)
}

func (cube Cube) volume() float64 {
	return math.Pow(cube.side, 3)
}

func (box Box) volume() float64 {
	return box.length * box.width * box.height
}

func (box Box) String() string {
	return fmt.Sprintf("The volume of the box with length:%v width:%v height:%v is %v", box.length, box.width, box.height, getVolume(box))
}

func (cube Cube) String() string {
	return fmt.Sprintf("The volume of the cube with side %v is %v", cube.side, getVolume(cube))
}

func (sphere Sphere) String() string {
	return fmt.Sprintf("The volume of the sphere with radius %v is %v", sphere.radius, getVolume(sphere))
}

func getVolume(structure Structure) float64 {
	return structure.volume()
}

func main() {
	cube := Cube{10}
	box := Box{10, 20, 20}
	sphere := Sphere{10}

	fmt.Println(cube)
	fmt.Println(sphere)
	fmt.Println(box)
}
