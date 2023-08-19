package main

import "fmt"


//user defined types

type entier int //a user defined types int

type Entier = int //alias of int

func main(){
	var x int =10
	var y Entier =12
	var res int
	res = somme(x,y)

	var resEn entier
	resEn= entier(somme(x,y))

	fmt.Printf("\n%T \n%T  \n%T \n%T \n%v", x, y, res, resEn, resEn)
}

func somme(a,b int) int{
	return a+b
}
func sum (a,b int)entier{
	return entier(a+b)
}
