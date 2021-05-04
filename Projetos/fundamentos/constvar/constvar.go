package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tiop float64 é subententida pelo compilador

	area := PI * math.Pow(raio, 2)

	fmt.Println("A área de circuferência é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = false, true

	fmt.Println(e, f)

	g, h, i := true, 1, "epa!"

	fmt.Println(g, h, i)

}
