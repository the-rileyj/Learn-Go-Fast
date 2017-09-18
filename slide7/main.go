package main

import "fmt"

func main() {
	var boolean bool = true
	var characters string = "ayy, yuh"
	var unsigned_integer uint = 1
	var integer int = -1
	var character rune = 'R'
	var float = 66.94
	fmt.Printf("%T: %t;\n", boolean, boolean)
	fmt.Printf("%T: %s;\n", characters, characters)
	fmt.Printf("%T: %d;\n", unsigned_integer, unsigned_integer)
	fmt.Printf("%T: %d;\n", integer, integer)
	fmt.Printf("%T: %c;\n", character, character)
	fmt.Printf("%T: %.1f", float, float)
}

/*
func main() {
	var boolean bool = true
	var characters string = "ayy, yuh"
	var unsigned_integer uint = 1
	var integer int = -1
	var character rune = 'R'
	var float = 66.94
	fmt.Printf("%T: %v;\n", boolean, boolean)
	fmt.Printf("%T: %v;\n", characters, characters)
	fmt.Printf("%T: %v;\n", unsigned_integer, unsigned_integer)
	fmt.Printf("%T: %v;\n", integer, integer)
	fmt.Printf("%T: %v;\n", character, character)
	fmt.Printf("%T: %.1v", float, float)
}*/
