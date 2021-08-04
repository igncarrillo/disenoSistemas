package handler

import (
	"helloWorld/service"
)

type Handler interface {
	Run()
}

type handler struct {
	service service.Service
}

func (h *handler) Run(){
	h.service.RunService()
}
func NewHandler(s service.Service) Handler{
	return &handler{s}
}
