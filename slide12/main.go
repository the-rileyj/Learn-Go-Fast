package main

import "fmt"

func multiReturn() (int, string) {
	return 9, "We want this :)"
}

func main() {
	_, want := multiReturn()
	for _, c := range want {
		fmt.Printf("%c\n", c)
	}
}

/*
package main

import "fmt"

func multiReturn() (int, string) {
	return 9, "We want this :)"
}

func main() {
	_, want := multiReturn()
	for i, _ := range want {
		fmt.Printf("%d\n", i)
	}
}

package main
import "fmt"

func main() {
	for range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		fmt.Println("I'm looping")
	}
}
*/
