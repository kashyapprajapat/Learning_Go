package main

import "fmt"

func main() {
	fmt.Println("If-else")

	age := 20
	if age>=22 {
		fmt.Println("Yes above 22")
	} else {
		fmt.Println("Below 22")
	}

	score := 700
    if score >=699 && score <=1000{
		fmt.Println("Valid score")
	}


	// go if feature
	if marks:= 500; marks >=499 {
		fmt.Println(marks)
	}

	//Go doesnot have ternary operator
}