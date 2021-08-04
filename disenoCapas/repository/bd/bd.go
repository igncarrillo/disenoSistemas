package bd

import (
	"fmt"
	repo "helloWorld/repository"
)


type repository struct {
	name string
}

func (r *repository) Escribir()  {
	fmt.Println("Estoy escribiendo en ",r.name)
}

func NewRepository(n string) repo.Repository {
	return &repository{n}
}