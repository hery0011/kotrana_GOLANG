package main

import "fmt"

func main() {

	/*for i := 0; i < 5; i++ {
		fmt.Println(i)
	}*/

	x := 0
	/*for x <= 5 {
		fmt.Println(x)
		x++
	}*/

	/*for {
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
