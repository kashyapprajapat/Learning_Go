package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	//Remember 
	// int -> 0 , string -> "", bool -> false 

	var nums [4] int 
	//array lenght
	fmt.Println(len(nums))

	//add
	nums[0] = 1
	fmt.Println(nums[0])

	//print whole array
	fmt.Println(nums)
  
	fmt.Println("============================")

	fmt.Println("bools Array")

	var kas[4] bool
	fmt.Println(kas)

	fmt.Println("============================")

	fmt.Println("String Array")

	var mak[4] string
	mak[1] = "kp"
	fmt.Println(mak)

	// You can direct assign  the values
    //nums := [3]int{10,20,30}

}
