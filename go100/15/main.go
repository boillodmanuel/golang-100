package main

import "fmt"

func pa(msg string, a *[4]int) {
	fmt.Printf("%v, array %p, len=%v cap=%v %v\n", msg, a, len(a), cap(a), *a)
}
func ps(msg string, a []int) {
	fmt.Printf("%v, slice %p, len=%v cap=%v %v\n", msg, a, len(a), cap(a), a)
}

func main() {
	var a = [4]int{0, 1 ,2 ,3}
	c := a[0: cap(a)]
	d := c

	pa("a", &a)
	ps("c", c)
	ps("d", d)
	fmt.Println()

	c[1] = -1
	pa("a", &a)
	ps("c", c)
	ps("d", d)
	fmt.Println()

	c = append(c, 4)
	c[2] = -2
	pa("a", &a)
	ps("c", c)
	ps("d", d)
	fmt.Println()

	c = append(c, 6)
	pa("a", &a)
	ps("c", c)
	ps("d", d)
	fmt.Println()

	c[1] = -1
	pa("a", &a)
	ps("c", c)
	ps("d", d)
}
