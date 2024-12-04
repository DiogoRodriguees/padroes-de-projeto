package publishers

import (
	"log"
	"observer/interfaces"
)

type Publisher struct {
	Observers []interfaces.ISubscriber
}

func New() *Publisher {
	return &Publisher{
		Observers: []interfaces.ISubscriber{},
	}
}

func (s *Publisher) Subscribe(obs interfaces.ISubscriber) {
	s.Observers = append(s.Observers, obs)
}

func (s *Publisher) UnSubscribe(obs interfaces.ISubscriber) {
	for i, v := range s.Observers {
		if v == obs {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
		}
	}
}

func (s *Publisher) UnSubscribeALl() {
	s.Observers = []interfaces.ISubscriber{}
}

func (s *Publisher) Notify(obs interfaces.ISubscriber) {
	log.Println("Notifying observer ", obs.GetID())
	obs.Update()
}

func (s *Publisher) NotifyAll() {
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

func (s *Publisher) NotifyByInterest(interesse string) {
	log.Println("Notifying all observers with interest in ", interesse)
	for _, obs := range s.Observers {
		if contains(obs.GetInterest(), interesse) {
			s.Notify(obs)
		}
	}
}
