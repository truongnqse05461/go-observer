package observer

import "testing"

func BenchmarkUpdate(b *testing.B) {
	d1 := NewDevice("sensor01", true)
	// d2 := NewDevice("hub01", true)

	client1 := NewClient("c1")
	client2 := NewClient("c2")
	client3 := NewClient("c3")
	client4 := NewClient("c4")
	client5 := NewClient("c5")
	client6 := NewClient("c6")
	client7 := NewClient("c7")
	client8 := NewClient("c8")
	client9 := NewClient("c9")

	d1.Subscribe(client1)
	d1.Subscribe(client2)
	d1.Subscribe(client3)
	d1.Subscribe(client4)
	d1.Subscribe(client5)
	d1.Subscribe(client6)
	d1.Subscribe(client7)
	d1.Subscribe(client8)
	d1.Subscribe(client9)
	// d2.Subscribe(client1)
	// d1.Subscribe(client2)

	for i := 0; i < b.N; i++ {
		d1.Update(i%2 == 0)
	}
}
