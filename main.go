package main

import "fmt"

func main() {
	// TODO: Add your code here
	fmt.Println("hello world")
	fmt.Println(People{
		Name: "cheng",
		Age:  18,
	})
}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
