package main

import "github.com/truongnqse05461/go-observer/observer"

func main() {
	d1 := observer.NewDevice("sensor01", true)
	d2 := observer.NewDevice("hub01", true)

	client1 := observer.NewClient("c1")
	client2 := observer.NewClient("c2")

	d1.Subscribe(client1)
	d2.Subscribe(client1)
	d1.Subscribe(client2)
	d1.Update(false)
	d2.Update(false)
	d2.Update(true)
	d1.Update(true)
}
