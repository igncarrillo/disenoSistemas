package main

import (
	"SingletonPattern/singleton"
	"fmt"
)

func main() {
	db := singleton.GetInstance("MySQL")
	fmt.Println(db.Getname())

	otherDB := singleton.GetInstance("otroNombre")
	fmt.Println(otherDB.Getname())

	//Vemos que sigue devolviendo la misma instancia
}
