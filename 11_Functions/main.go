package main


import "fmt"

func add(a int,b int) int {
	return a+b
}

//version2
func addv2(a,b int) int {
	return a+b
}


func getlang()(string,string,bool){
	return "Go","Js",true
}


func proceesit(fn func(a int) int){
	fn(7)
}


// Go main Feature in function that 
// Go can return multiple value in the function
func main() {
	fmt.Println("functions")
    fmt.Println(add(5,2))
    fmt.Println(addv2(3,4))

	//Remebr in go we rteun two value
	// first is  actual fun want tot return
	// second one i sth error
	fmt.Println(getlang())

	fn := func(a int) int{
		return a*a
	}
	proceesit(fn)
	fmt.Println()


}