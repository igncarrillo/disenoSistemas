package creator

import (
	"FactoryPattern/product"
	"errors"
)

type Logistica interface {
	AsignarTransporte() product.ITransporte
}

type LogisticaMaritima struct {

}

type LogisticaArea struct {

}

func (lm *LogisticaMaritima) AsignarTransporte() product.ITransporte{
	return &product.Barco{}
}

func (la *LogisticaArea ) AsignarTransporte() product.ITransporte{
	return &product.Avion{}
}

func GetPlatform(s string) (Logistica, error){
	switch s {
	case "maritima":
		return &LogisticaMaritima{}, nil
	case "aerea":
		return &LogisticaArea{}, nil
	default:
		return nil, errors.New("Tipo de plataforma invalida")
	}
}
