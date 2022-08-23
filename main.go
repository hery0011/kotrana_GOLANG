package main

import (
	"fmt"
)

func main() {
	var nbr1 int = 40
	var nbr2 int = 5
	//var result int = nbr1 / nbr2

	var nom string = "hery"
	var age int = 23
	//nom := "hery"

	fmt.Println("Bonjour, mon Nom est", nom, "et j'ai", age, "ans")
	//fmt.Println("le resultat est", result)

	fmt.Println("===========")
	if result := nbr1 / nbr2; result == 8 {
		fmt.Println("mitovy")
	} else {
		fmt.Println("tsy mitovy")
	}
}
