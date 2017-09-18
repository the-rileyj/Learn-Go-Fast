package main

import "fmt"

func main() {
	exMap := make(map[string]int)
	exMap["ayy"] = 9
	fmt.Println(exMap["ayy"], exMap["yuh"])
	liesMap := map[string]bool{"RJ is neato": true, "Tom sucks at C": true,
		"Cody will give RJ extra credit": false}
	fmt.Println("Here is the 'liesMap' printed as a 'truthMap':")
	for key, value := range liesMap {
		fmt.Printf("%s: %t\n", key, !value)
	}
}
