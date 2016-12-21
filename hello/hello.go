package main

import "fmt"

func getName() string {
	return "Go"
}

func main () {
	name := getName()
	fmt.Println("This is a first ", name)
}