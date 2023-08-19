package main

import "fmt"

type void func()

func main() {
	var f void
	process(func(){})
	fmt.Printf("\n %T",f)
}

func process(f void){
	f()
}
