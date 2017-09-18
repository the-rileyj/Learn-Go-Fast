package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("I'm looping")
	}
}

/*
func main() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println("I'm looping")
	}
}

func main() {
	i := 0
	for i < 10 {
		fmt.Println("I'm looping")
		i++
	}
}

func main() {
	i := 0
	for {
		fmt.Println("I'm looping")
		i++
		if i >= 10 {
			break
		}
	}
}
*/
