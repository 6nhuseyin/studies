package main

import "fmt"

func makeMeOlder(age int) {
	age += 5
  }
  
  func main() {
	myAge := 10
	makeMeOlder(myAge)
	fmt.Println(myAge)
	// myAge is still 10
  }