package main

import "fmt"

// slice -> Dynamic Array

func main(){

	// Uninitialized slice is nill
	var nums[] int

	fmt.Println(nums)
    fmt.Println(nums==nil)
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	nums = append(nums, 40)

    fmt.Println(nums)

	//Slice operator
	var name = []int{10,20,30,40,50}
	fmt.Println(name[0:2])

}