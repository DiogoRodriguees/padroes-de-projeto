package interfaces

type ISubject interface {
	Subscribe()
	UnSubscribe()
	UnSubscribeALl()
	Notify()
	NotifyAll()
}
