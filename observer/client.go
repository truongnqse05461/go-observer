package observer

import (
	"fmt"
)

type client struct {
	id string
}

// GetID implements Observer
func (c *client) GetID() string {
	return c.id
}

// Update implements Observer
func (c *client) Update(m map[string]interface{}) {
	var status string
	switch m["status"].(bool) {
	case true:
		status = "online"
	case false:
		status = "offline"
	}
	fmt.Printf("[%s] - device [%s] change status to [%s]\n", c.id, m["name"], status)
}

func NewClient(id string) Observer {
	return &client{id: id}
}

var _ Observer = (*client)(nil)
