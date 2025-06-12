package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://kashyapprajapati.netlify.app/"

func main() {
	fmt.Println("Kp Portfolio web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type is : %T\n", response)
	defer response.Body.Close() // caller responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
