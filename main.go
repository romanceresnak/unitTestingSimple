package main

import "fmt"

func HelloGo() string{
	return "Welcome to go language"
}

func main(){
	output := HelloGo()
	fmt.Println(output)
}
