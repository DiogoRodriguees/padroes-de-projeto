package infra

import (
	"fmt"
	"observer/interfaces"
)

type Subject struct {
	Observers []interfaces.IObserver
}

func NewSubject() *Subject {
	return &Subject{
		Observers: []interfaces.IObserver{},
	}
}

func (s *Subject) Subscribe(obs interfaces.IObserver) {
	s.Observers = append(s.Observers, obs)
}

func (s *Subject) UnSubscribe(obs interfaces.IObserver) {
	for i, v := range s.Observers {
		if v == obs {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
		}
	}
}

func (s *Subject) UnSubscribeALl() {
	s.Observers = []interfaces.IObserver{}
}

func (s *Subject) Notify(obs interfaces.IObserver) {
	fmt.Println("Notifying observer ", obs.GetID())
	obs.Update()

}

func (s *Subject) NotifyAll() {
	for _, obs := range s.Observers {
		s.Notify(obs)
	}
}
