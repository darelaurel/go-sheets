package main

import "fmt"

func main() {
	h :=make(map[string]int)
	h["one"] = 17

	//know element existence with found var

	res, found :=h["one"]

	fmt.Print(res,found)

	fmt.Printf("\n%v", h)
}
