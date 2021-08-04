package main

import (
	"fmt"
	"helloWorld/handler"
	"helloWorld/repository/bd"
	"helloWorld/service"
)

func main() {

	fmt.Println("Hello World !")

	bdr := bd.NewRepository("MYSQL")
	//api := microservice.NewRepository("coreApplication")
	s:= service.NewService(bdr)
	h:= handler.NewHandler(s)

	h.Run()
}
