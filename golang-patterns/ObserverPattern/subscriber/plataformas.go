package subscriber

import "fmt"

type Subscriber interface {
	Update(s string)
}

type slack struct {

}

type email struct {

}

func (e *email) Update(s string)  {
	fmt.Println("Enviando por email: ",s)
}

func NewEmail() Subscriber{
	return &email{}
}

func (sl *slack) Update(s string)  {
	fmt.Println("Enviando por slack: ",s)
}

func NewSlack() Subscriber{
	return &slack{}
}