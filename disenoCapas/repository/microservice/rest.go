package microservice

import (
	"fmt"
	repo "helloWorld/repository"
)

type repository struct {
	micros string
}

func (r *repository) Escribir()  {
	fmt.Println("Haciendo una pegada al microservicio: ",r.micros)
}

func NewRepository(api string) repo.Repository {
	return &repository{api}
}

