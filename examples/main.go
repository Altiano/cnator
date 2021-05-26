package main

import (
	"time"

	"gitlab.com/altiano/cnator"
)

func main() {

	cnator := cnator.New()
	channels := createChannels()

	// provide channel reference to each publisher
	producerA := NewProducerA(channels.chanA)
	producerB := NewProducerB(channels.chanB)
	producerC := NewProducerC(channels.chanC)

	// using cnator to subscribe to those channel events
	subscriberA := subscriberA{}
	subscriberB := subscriberB{}
	subscriberC := subscriberC{}
	cnator.Subscribe(channels.chanA, subscriberA.receiveChanA)
	cnator.Subscribe(channels.chanB, subscriberB.receiveChanB)
	cnator.Subscribe(channels.chanC, subscriberC.receiveChanC)

	// Need to call this after you have orchestrade your subscriptions design.
	// It will perform runtime check.
	// Panics if the subscriber parameter mismatch with the channel it subscribes to
	// Or when the one of the channel is not initialized (nil, it happens i know)
	cnator.Serve()

	// Simulate running program
	for i := 0; i < 3; i++ {
		producerA.publish()
		producerB.publish()
		producerC.publish()
	}

	<-time.After(2 * time.Second)
}
