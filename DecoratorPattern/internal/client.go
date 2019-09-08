package internal

import "time"

// Performer is something that performs an action which usually takes some amount of time such as a network call any blocking function.
type Performer interface {
	PerformAction()
}

// Client is an implementation of Performer. timeTaken is used to control the amount of time it takes to run the function
type client struct {
	timeTaken int
}

// Returns an Instance of Performer
func NewClient(time int) Performer{
	return &client{timeTaken:time}
}

func(c *client) PerformAction() {
	time.Sleep(time.Duration(c.timeTaken)*time.Second)
}