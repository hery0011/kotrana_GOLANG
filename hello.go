package main

import "fmt"

func helloHery() {
	fmt.Println("hery Rasolonjatovo")
}

func keyVal() {
	valKey := map[string]int{
		"janvier": 10,
		"fevrier": 20,
		"mars":    30,
	}
	fmt.Println(valKey["mars"])
}
