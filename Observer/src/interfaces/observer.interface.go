package interfaces

type IObserver interface {
	Update()
	GetID() int
	GetInterest() []string
}
