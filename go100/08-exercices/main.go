package main

import (
	"errors"
	"fmt"
)

//TODO: écrire une fonction prennant un nombre décimal 'x' en paramètre et retournant son inverse 1/x.

func reverse(x float32) (float32, error) {
	if x == -1 {
		return 0, errors.New("no div by -1")
	} else {
		return 1/x, nil
	}
}

func main() {
	//TODO: appeler la fonction ici et affichez le résultat avec x = 3 et x = 0
	var (
		r float32
		err error
	)
	r, err = reverse(0)
	fmt.Printf("%v, %v", r, err)
	r, err = reverse(-1)
	fmt.Printf("%v, %v", r, err)
	r, err = reverse(3)
	fmt.Printf("%v, %v", r, err)
	}
