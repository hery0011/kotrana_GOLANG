package main

import "fmt"

var x int = 0

func direBonjour(nom string) {
	println("bonjour", nom)
}

func kotranaCondition(nbr1, nbr2 int) string {

	if result := nbr1 / nbr2; result == 8 {
		return "mitovy"
	} else {
		return "tsy mitovy"
	}
}

func kotranaBoucle() {
	/*for i := 0; i < 5; i++ {
		fmt.Println(i)
	}


	for x <= 5 {
		fmt.Println(x)
		x++
	}

	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}*/

	for ; x <= 10; x++ {
		if x%2 == 1 {
			continue
		}
		fmt.Println(x)
	}

}

func table() int {
	/*declaration array puis affectation valeur*/

	/*var tableau [2]int
	tableau[0] = 10
	tableau[1] = 20

	fmt.Println(tableau)*/

	//declaration avec affectation valeur sans preciser le nbr array
	tabl := [...]int{10, 20, 30}
	return tabl[1]
}

func main() {

	//appel fonction dire bonjour
	fmt.Println("=================hello")
	direBonjour("hery")

	//appel fonction condition
	fmt.Println("=================condition")
	resultCondition := kotranaCondition(40, 5)
	fmt.Println(resultCondition)

	//affichage valeur tableau
	fmt.Println("=================tableau")
	valTabl := table()
	fmt.Println(valTabl)

	//appel function kotranaBoucle
	fmt.Println("=================boucle")
	kotranaBoucle()

	//appel function dans autre fichier
	fmt.Println("=================appel function dans un autre package")
	helloHery()
	keyVal()

}
