package main

import (
	"observer/infra"
)

type IObserver interface {
	Update()
	GetID() int
}

type ISubject interface {
	Subscribe()
	UnSubscribe()
	UnSubscribeALl()
	Notify()
	NotifyAll()
}

func main() {
	observer1 := &infra.Observer{ID: 1}
	observer2 := &infra.Observer{ID: 2}
	observer3 := &infra.Observer{ID: 3}

	subject1 := infra.NewSubject()
	subject1.Subscribe(observer1)
	subject1.Subscribe(observer2)
	subject1.Subscribe(observer3)
	subject1.NotifyAll()
	subject1.Notify(observer3)
}
