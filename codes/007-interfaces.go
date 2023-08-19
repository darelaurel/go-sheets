/*
*Interface applicable aux structs et user defined types
*/
package main

import "fmt"

type printable interface{
	print() string
}

type todoList []*int

type user struct {
	id    int
	name  string
}

func (tl todoList) print() string{
	var str string
	str =fmt.Sprintf("Taille des éléments fait %d",len(tl))
	return str
}

func (u user) print() string{
	return fmt.Sprintf("\n PRINT DE USER %v\n", u)
}

func details(t printable){
	fmt.Println(t.print())
}

func main() {
	u := user{id: 1, name: "Chiba"}

	fmt.Printf("\n PREV U VALUE %v %p\n",ss, &ss)

	details(u)

}
