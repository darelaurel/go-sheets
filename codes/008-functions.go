package main

import (
	"errors"
	"fmt"
)

//currying
/**
** Like closure
 */
func makeAdd(a int) func(int) int {
	return func(b int) int {
		fmt.Printf("\n%d", b)
		return a + b
	}
}

/**
*Fn variadics
*recoit plusieurs params
*/

func suum(nbs ...int) int{
	var tmp int
	for _,x:=range nbs{
		tmp=tmp+x
	}
	return tmp	
}

//multiple returns
func divide(a,b int) (int,error){
	if(b !=0){
		return (a/b),nil
	}	
	return 0, errors.New("Rien a afficher")
}

func main() {
	s := makeAdd(12)(15)
	//s := makeAdd(12)
	//res := s(15)
	sc :=suum(2,5,8)
	dc,err:=divide(12,0)

	fmt.Printf("\n%v \n%v \n%v \n%v", s, sc, dc,err)
}
