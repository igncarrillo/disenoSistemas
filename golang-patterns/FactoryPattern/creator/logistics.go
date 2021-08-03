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

func GetPlatform(s string) (Logistica, error) {
	switch s {
	case "maritima":
		return &logisticaMaritima{}, nil
	case "aerea":
		return &logisticaArea{}, nil
	default:
		return nil, errors.New("Tipo de plataforma invalida")
	}
}
