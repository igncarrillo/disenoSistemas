package main

import (
	"ObserverPattern/publisher"
	"ObserverPattern/subscriber"
)

func main() {
	pb := publisher.NewMensaje()
	sl := subscriber.NewSlack()
	em := subscriber.NewEmail()

	pb.AddSubscriber("slack",sl)
	pb.AddSubscriber("email",em)

	pb.Notificar("Nuevo mensaje")

	pb.RemoveSubscriber("email")
	pb.Notificar("Mensaje solo para slack")
}
