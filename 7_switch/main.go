package main

import "fmt"

func main() {
	fmt.Println("Switech statements")

	j := 7
	switch j {
	case 2 : 
	      fmt.Println("Numebr is 2")
		//   break  -- go will hanlde by deafult it
	case 7 : 
	      fmt.Println("Numebr is 7")
		//   break
	default : 
	      fmt.Println("Numebr is not a number")
		  
	}
}