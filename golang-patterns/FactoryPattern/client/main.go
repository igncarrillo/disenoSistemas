package main

import (
	"FactoryPattern/creator"
	"fmt"
	"os"
)

func main() {
	//Levantamos de la config el tipo de plataforma
	p, err := creator.GetPlatform("aerea")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	p.RealizarEnvio()

}
