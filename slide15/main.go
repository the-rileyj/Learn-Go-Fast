package main

import "fmt"

func main() {
	var i int
	ptr := &i
	fmt.Println(i)
	*ptr = 9
	fmt.Println(i)
}
