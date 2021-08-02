package product

import "fmt"

type ITransporte interface {
	RealizarEnvio()
}

type Avion struct {
}

type Barco struct {
}

func (a *Avion) RealizarEnvio()  {
	fmt.Println("Realizando envio en avion")
}

func (b *Barco) RealizarEnvio()  {
	fmt.Println("Realizando envio en barco")
}