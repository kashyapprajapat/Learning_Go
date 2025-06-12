package main

import "fmt"

// Maps -->key value
func main(){
   

	kmap :=map[string]int{"Age":22,"numbercode":123}
	fmt.Println(kmap)
	//Create map--> map[key]vlaue both type 
	m := make(map[string]string)
	fmt.Println(m)


	//Setting an element
	m["name"]="Kashyap"
	m["Course"]="MCA"
	fmt.Println(m)

	//Get an element
	fmt.Println(m["name"])

    fmt.Println(m["neme"]) // Empty value if keyvalue not found
	
	//Length of map
	fmt.Println(len(m))  //2 
	

}