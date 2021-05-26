package main

type (
	channels struct {
		chanA chan struct{}
		chanB chan string
		chanC chan Person
	}

	Person struct {
		name  string
		email string
	}
)

func createChannels() channels {
	return channels{
		chanA: make(chan struct{}),
		chanB: make(chan string),
		chanC: make(chan Person),
	}
}
