package main

import "fmt"

func main () {
	name := "Codecademy"
	fmt.Printf("Address of %v is %v", name, &name)
	// Address of Codecademy is 0xc000010210
	fmt.Println(" ")
}