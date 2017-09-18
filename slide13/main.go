package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s1 []int = primes[1:4]
	var s2 []int = primes[:4] //[0:4]
	var s3 []int = primes[1:] //[1:6]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

/*
package main

import "fmt"

func main() {
	meaning_cool := []string{"neato", "radical"}
	fmt.Println(meaning_cool)
	meaning_cool = append(meaning_cool, "legit")
	fmt.Println(meaning_cool)
}
*/
