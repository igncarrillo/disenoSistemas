package creator

import (
	"FactoryPattern/product"
	"errors"
)

type Logistica interface {
	AsignarTransporte() product.ITransporte
}

type logisticaMaritima struct {
}

type logisticaArea struct {
}

func (lm *logisticaMaritima) AsignarTransporte() product.ITransporte {
	return &product.Barco{}
}

func (la *logisticaArea) AsignarTransporte() product.ITransporte {
	return &product.Avion{}
}

func GetPlatform(s string) (product.ITransporte, error) {
	switch s {
	case "maritima":
		m := &logisticaMaritima{}
		return m.AsignarTransporte(), nil
	case "aerea":
		a:= &logisticaArea{}
		return a.AsignarTransporte(), nil
	default:
		return nil, errors.New("Tipo de plataforma invalida")
	}
}
