package main

import (
	"observer/infra/publishers"
	"observer/infra/subscribers"
)

func main() {
	subscriber1 := subscribers.New(1, []string{"Esportes", "Noticias"})
	subscriber2 := subscribers.New(2, []string{"Noticias"})
	subscriber3 := subscribers.New(3, []string{"Esportes"})

	publisher := publishers.New()
	publisher.Subscribe(subscriber1)
	publisher.Subscribe(subscriber2)
	publisher.Subscribe(subscriber3)
	publisher.NotifyAll()                  // Notifica todos os observers (1, 2 e 3)
	publisher.NotifyByInterest("Esportes") // Notifica assinante com id 1 e 3
	publisher.NotifyByInterest("Noticias") // Notifica assinante com id 1 e 2
}
