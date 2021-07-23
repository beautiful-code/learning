# Structure Interface

Given Structure interface 
	type Structure interface {
		volume() float64
	}

Implement Structure interface for Cube, Box and Sphere structs.
Add a getVolume() function that accepts Structure interface and computes volume of the input.

Additionally, each of these structs should also implement fmt.Stringer interface
so that output of getVolume() is printed out with struct specific information.
Eg: For Sphere input, 
Output should be "The volume of the sphere with radius 5 is 523.59878"
 
## Coding the solution

Look for main.go file and place your solution over there.