package service

import (
	"helloWorld/repository"
)

type Service interface {
	RunService()
}

type service struct {
	repo repository.Repository
}

func (s *service) RunService()  {
	s.repo.Escribir()
}
func NewService(r repository.Repository) Service{
	return &service{r}
}