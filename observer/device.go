package observer

type device struct {
	observers []Observer
	name      string
	status    bool
}

// NotifyAll implements Publisher
func (d *device) NotifyAll() {
	for _, obs := range d.observers {
		obs.Update(map[string]interface{}{
			"name":   d.name,
			"status": d.status,
		})
	}
}

// Subscribe implements Publisher
func (d *device) Subscribe(observer Observer) {
	d.observers = append(d.observers, observer)
}

// Unsubscribe implements Publisher
func (d *device) Unsubscribe(observer Observer) {
	d.observers = remove(d.observers, observer)
}

// Update implements Publisher
func (d *device) Update(status bool) {
	d.status = status
	d.NotifyAll()
}

func remove(observers []Observer, removeObs Observer) []Observer {
	obsLength := len(observers)
	for i, obs := range observers {
		if obs.GetID() == removeObs.GetID() {
			observers[i], observers[obsLength-1] = observers[obsLength-1], observers[i]
			return observers[:obsLength-1]
		}
	}
	return observers
}

func NewDevice(name string, status bool) Publisher {
	return &device{name: name, status: status}
}

var _ Publisher = (*device)(nil)
