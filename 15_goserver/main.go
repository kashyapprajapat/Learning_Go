package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Weelcome to golang server")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "https://reviewverse.onrender.com/users"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Print(response)
	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	content ,err := ioutil.ReadAll(response.Body)

	fmt.Println(content) // <-- by default it is in byte format
	fmt.Println(string(content))
}