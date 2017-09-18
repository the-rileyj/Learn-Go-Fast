package main

import "fmt"

func main() {
	fmt.Printf("That doggo is a ")
	if doggo := 1; doggo <= 0 {
		fmt.Println("yipper")
	} else if doggo == 1 {
		fmt.Println("pupper")
	} else {
		fmt.Println("woofer")
	}
}

/*
package main
import "fmt"

func main() {
	fmt.Printf("That doggo is a ")
	switch doggo := -1; doggo {
	case -1:
		fmt.Printf("yapper and a ")
		doggo += 1
		fallthrough
	case 0:
		fmt.Println("yipper")
	case 1:
		fmt.Println("woofer")
	default:
		fmt.Println("doggo")
	}
}
*/
