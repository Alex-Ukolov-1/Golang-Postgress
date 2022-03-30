package main

import(
	"log"
	"fmt"
)


func main(){
	e:=Estudiane{
		Name:"Alex",
		Age:20,
		Active: true,
	}

	err:=Clear(e)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Creado exitsamente")
}