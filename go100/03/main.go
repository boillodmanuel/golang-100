//Identifiants exportés en dehors de leur package
package main

import (
	"fmt"
	"math"
)

var jeSuisInvisibleEnDehorsDuPackage int

var JeSuisVisibleEnDehorsDuPackage int

func main() {
	fmt.Println(math.Pi) //cannot refer to unexported name math.pi
}
