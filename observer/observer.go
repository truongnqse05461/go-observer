package observer

type Observer interface {
	Update(map[string]interface{})
	GetID() string
}

type Publisher interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Update(bool)
	NotifyAll()
}
