package main

import "fmt"

//arrays : taille fixe

//slices : taille dynamique

//append is like push can't work array only slice types

func main() {
	var arr [5]int
	var slice []int
	for i := 0; i < 5; i++ {
		arr[i] += i + 1
	}

	for i := 0; i < 9; i++ {
		slice = append(slice, i+1)
	}

	//read arr or slices element
	for id,x :=range arr{
		fmt.Println(id,x)
	}

	fmt.Printf("\n%v \n%v", arr, slice)
}
