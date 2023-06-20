package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`Hello "Huseyin"`)
	fmt.Println("This string\nspans multiple\nlines.")
	fmt.Println(`Sammy says,\"The balloon\'s color is red.\"`)
	fmt.Println("1.\tShark\n2.\tShrimp\n10.\tSquid")

	fmt.Println("This string is on\n"+
	"multiple lines\n" +
	"within three single\n" +
	"quotes on either side.")
	ss := "Sammy Shaaaaaaaaaark"
	fmt.Println(strings.ToLower(ss))
	fmt.Println(strings.HasPrefix(ss ,"Sa"))
	fmt.Println(strings.HasSuffix(ss, "ark"))
	fmt.Println(strings.Count(ss, "a"))
}	