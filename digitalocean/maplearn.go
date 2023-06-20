package main  
import ("fmt")  
 
func main() {
	
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)


}
/*
func main() {  
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}  
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}  
  
  fmt.Printf("a\t%v\n", a)  
  fmt.Printf("b\t%v\t\n", b)
}
*/