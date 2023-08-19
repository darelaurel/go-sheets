package main

import "fmt"

/**
*Les embedded structs doivent toujours etre passé par référence (le pointeur)
*Les operations sur les embedded structs doivent etre faits par reference comme addTodo
* de sorte que si l'embedded struct change embedding struct (struct dans laquelle une autre
est incorporé change automatiquement case de t.toggle)
**/

type todo struct {
	name string
	done bool
}

func (t *todo) toggle() {
	t.done = !t.done
}

//embedded structs
/**
* structs incorporés dans une autre
**/

type user struct {
	id    int
	name  string
	todos []*todo
}

//Function receiver is bind fn on struct element like addTodo

func (u *user) addTodo(t *todo) {
	u.todos = append(u.todos, t)
}

func (u user) show() {
	for id, x := range u.todos {
		fmt.Println(id, x)
	}
	fmt.Printf("\n UTODOS %v  %p\n", u.todos, u.todos)
	// fmt.Printf("\nVALEUR DE MON U SANS POINTEUR %v\n", u)
	// fmt.Printf("POINTEUR DE MON U SANS POINTEUR %p\n", &u)
}

var s = new(user)

func main() {
	u := user{id: 1, name: "Chiba"}
	t := todo{name: "Lire Go course", done: false}

	fmt.Printf("\n PREV U VALUE %v %p\n", u, &u)

	fmt.Print("------------------------\n\n")

	t.toggle()
	u.show()

	u.addTodo(&t)

	t.toggle()
	u.show()

	fmt.Printf("\nEND U VALUE %v %p\n", u, &u)
	//fmt.Printf("\n%v\n \n%p\n \n%v\n", u, s, t)
}
