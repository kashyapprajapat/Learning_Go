package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doing task",id)
}


func main() {
	fmt.Println("Goruines means mutilthreading conccurannncy--High cpu intensive task perfoem based on the cpu core")

	for i:=1;i<=10;i++{
		go task(i)
	}

	time.Sleep(time.Second * 3)
}