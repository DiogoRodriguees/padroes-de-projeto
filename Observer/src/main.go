package main

import (
	"observer/infra/observers"
	"observer/infra/subjects"
)

func main() {
	observer1 := observers.New(1, []string{"Esportes", "Noticias"})
	observer2 := observers.New(2, []string{"Noticias"})
	observer3 := observers.New(3, []string{"Esportes"})

	subject := subjects.New()
	subject.Subscribe(observer1)
	subject.Subscribe(observer2)
	subject.Subscribe(observer3)
	subject.NotifyAll()                  // Notifica todos os observers (1, 2 e 3)
	subject.NotifyByInterest("Esportes") // Notifica observer com id 1 e 3
	subject.NotifyByInterest("Noticias") // Notifica observer com id 1 e 2
}
