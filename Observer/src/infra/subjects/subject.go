package subjects

import (
	"log"
	"observer/interfaces"
)

type Subject struct {
	Observers []interfaces.IObserver
}

func New() *Subject {
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
	log.Println("Notifying observer ", obs.GetID())
	obs.Update()

}

func (s *Subject) NotifyAll() {
	log.Println("Notifying all observers")
	for _, obs := range s.Observers {
		s.Notify(obs)
	}
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func (s *Subject) NotifyByInterest(interesse string) {
	log.Println("Notifying all observers with interest in ", interesse)
	for _, obs := range s.Observers {
		if contains(obs.GetInterest(), interesse) {
			s.Notify(obs)
		}
	}
}
