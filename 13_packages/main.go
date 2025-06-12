package main

import (
	"fmt"
	"image/color"
	"kashyap/auth"
	"kashyap/user"
)

//in package just remmber you want export one fun to to am=nother object so
// use cappital latter in the function
// if you want to make public anduse outisde so Every first allter capital
// if you want to make it private then make a small so
// that pther than that package no cna use that

// in node js we have npm for packages samee way in this we ahve
// go packages -->https://pkg.go.dev/
// to install go package use this coomand go get <packagename>

func main() {
	auth.Login("kp","K@1411p")
	name := user.UserName()
	fmt.Println(name)

	color.Black(name)
}