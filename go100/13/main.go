package main

import "fmt"

type Vertex struct {
	X, Y int
	Z *Vertex
}
type A struct {
	X, Y int
}
type B struct {
	X, Y int
	A A
}

var (
	v1 = Vertex{1, 2, nil} // has type Vertex
	v2 = Vertex{X: 1, Z: &v1} // Y:0 is implicit
	v3 = v1     // X:0 and Y:0
)

func main() {
	fmt.Println(v1, v2, v3)
	v1.X = 12
	fmt.Println(v1, v2, v3)

	a := A{1, 2}
	b := B{3, 4, a}
	fmt.Println(a, b)
}
