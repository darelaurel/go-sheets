package main

import "fmt"

//Initialiazition fn

func init() {
	fmt.Println("init")
}

type user struct {
	id   int
	name string
}

func main() {
	var s string
	var i int = 12
	var f float64
	var p *int
	var a []int

	p = &i

	u := user{id: 1, name: "CHiba"}
	name := "Darel"

	fmt.Printf("%s \n%d \n%f \n%p \n%v \n%d \n%s \n%v", s, i, f, p, a, *p, name, u)
}
