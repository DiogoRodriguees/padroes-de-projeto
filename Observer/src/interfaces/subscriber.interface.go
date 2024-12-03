package interfaces

type ISubscriber interface {
	Update()
	GetID() int
	GetInterest() []string
}
