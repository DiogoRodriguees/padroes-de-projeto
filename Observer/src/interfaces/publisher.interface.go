package interfaces

type IPublisher interface {
	Subscribe()
	UnSubscribe()
	UnSubscribeALl()
	Notify()
	NotifyAll()
}
