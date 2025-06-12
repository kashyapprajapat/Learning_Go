package main


import "fmt"


var myage =22

func main(){
	fmt.Println("Constant")
	const name = "kashyap"
	fmt.Println(name)
	fmt.Print(myage)

	const ( 
		port = 5000
		host = "loclahot"
    )

	fmt.Println(port,host)
}