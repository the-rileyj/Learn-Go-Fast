package main

import "fmt"

func main() {
	t := ""
	var r string
	var a = ""
	var f string = ""
	fmt.Println(t, r, a, f)
}

/*
package main
import "fmt"

var global = "I'm accessable anywhere in this program

func main() {
	local_to_main := "I'm only accessable in the main function"
	if true {
		local_to_if := "I'm only accessable in the if statement"
	}
}
*/
