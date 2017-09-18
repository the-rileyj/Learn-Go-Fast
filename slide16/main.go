package main

import "fmt"

type person struct {
	shirt     string
	underwear string
	pants     string
	socks     string
	shoes     string
}

func main() {
	RJ := person{"Neon", "Wouldn't you like to know", "Khaki", "Black", "White"}
	fmt.Printf("%+v\n", RJ)
}

/*package main
s
import "fmt"

type person struct {
	shirt     string
	underwear string
	pants     string
	socks     string
	shoes     string
}

func main() {
	RJ := person{shirt:"Neon", underwear:"Wouldn't you like to know", pants:"Khaki", socks:"Black", shoes:"White"}
	fmt.Printf("%+v\n", RJ)
}*/

/*package main

import "fmt"

type person struct {
	shirt     string
	underwear string
	pants     string
	socks     string
	shoes     string
}

func main() {
	RJ := person{}
	RJ.shirt = "Neon"
	RJ.underwear = "Wouldn't you like to know"
	RJ.pants = "Khaki"
	RJ.socks = "Black"
	RJ.shoes = "White"
	fmt.Printf("%+v\n", RJ)
}
*/
