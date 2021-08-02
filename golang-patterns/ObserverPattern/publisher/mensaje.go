package publisher

import "ObserverPattern/subscriber"

type Publisher interface {
	AddSubscriber(n string, s subscriber.Subscriber)
	RemoveSubscriber(n string)
	Notificar(mj string)
}
type mensaje struct {
	msg string
	subscribers map[string]subscriber.Subscriber
}

func (m *mensaje) AddSubscriber(n string, s subscriber.Subscriber) {
	if m.subscribers==nil{
		m.subscribers=make(map[string]subscriber.Subscriber)
	}
	m.subscribers[n]=s
}

func (m *mensaje) RemoveSubscriber(n string)  {
	delete(m.subscribers, n)
}

func (m *mensaje) Notificar(mj string)  {
	m.msg=mj
	for _,s := range m.subscribers{
		s.Update(mj)
	}
}

func NewMensaje() Publisher{
	return &mensaje{}
}
