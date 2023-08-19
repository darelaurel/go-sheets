package main

import (
	"fmt"
	"runtime/debug"
)

/*panic fait crasher l'appli
use case : chargement de var non trouvé pour demarrer l'appli
use case : element important pour run l'application
use case : a ne pas utiliser pour de simples var
*/

func boum() {
	panic("Oups, souci")
}
func main() {
	/**
		*Fn executé a la fin de l'execution de main (son contexte) en stack sil a
		*plusieurs defer
		*use case : faire du nettoyage, clean pointeur, var
	**/
	defer func() {
		fmt.Println("First defer")
	}()
	defer func() {
		/**
			*recuperer les elements d'un panic
			*tracker les panics
		**/
		res := recover()
		if res != nil {
			//fmt.Println("Panic",res)
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	boum()
}
