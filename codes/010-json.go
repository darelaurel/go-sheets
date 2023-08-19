package main

import (
	"fmt"
	"encoding/json"
)

type user struct{
	Name string `json:"user_name"`
	Age int `json:age`
	Groupe string `json:"groupe"`
}

func main(){
	u1 := user{Name:"James",Age:12, Groupe:"Demo"}

	/*
		return byte_slice of struct needs to be string before
		return json element
	*/
	res,err :=json.Marshal(u1)
	if(err !=nil){
		fmt.Println(err)
	}
	json_data :=string(res)
	fmt.Println(json_data)

	u2 :=user{}
	//Unmarshal take byte_slice and interface
	err = json.Unmarshal(res, &u2 )
	if(err !=nil){
		fmt.Println(err)
	}
	fmt.Println(u2)
}