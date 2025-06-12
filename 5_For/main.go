package main

import "fmt"


//For -> only construct in go for looping
func main() {
	fmt.Print("For Loop")

	//while loop
	i := 1
	for i<=10 {
		fmt.Println(i)
		i = i +1 
	}

	//classic for lop
	for j:=1;j<=5;j++ {
		fmt.Print(j)
		if j== 3 {
           continue
		}	
	}

	//Range
	for  k := range 7{
		fmt.Println(k)
	}

}