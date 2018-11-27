package main

import (
	"fmt"
)

//TODO: écrire une fonction qui prend en paramètre un entier et un nombre max d'itérations à effectuer, et qui effectue l'algorithme de
//la conjecture de Syracuse (voir README.md).
//Elle devra retourner un booléen indiquant si le nombre 1 a été atteint, suivi du nombre d'itérations qui ont été effectuées.
func syracuse(value int, maxIterationCount int) (reached bool, iterationCount int) {
	var i int;

	for i = 0; value != 1 && i < maxIterationCount; i++ {
		//log.Println(value, i)
		if value % 2 == 0 {
			value = value / 2
		} else {
			value = 3 * value + 1
		}
	}
	return value == 1, i
}

func print(value int, maxIterationCount int) {
	reached, iterationCount := syracuse(value, maxIterationCount)
	if reached {
		fmt.Printf("L'algorithme a permis d'atteindre le nombre 1 à partir du nombre %d en %d itérations\n", value, iterationCount)
	} else {
		fmt.Printf("L'algorithme n'a pas permis d'atteindre le nombre 1 à partir du nombre %d en moins de %d itérations\n", value, maxIterationCount)
	}

}

func main() {
	print (98375, 1000)
	print (15, 1000)
}
