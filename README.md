# Intro
A ulitity package to enable Channel-based subscriptions in your golang code.

It enables "event-driven" approach in your application.

It has runtime check including when the subscribers provide invalid parameter. 


# Usage
```
go get gitlab.com/altiano/cnator
```

```
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
```

# TODO
- [ ] Tracing visibility of every calls happen in subscriptions. Useful for debugging.
- [ ] Supports channels to shared by multiple `Cnator` instances


# Project Icon

Photo by <a href="https://unsplash.com/@gabiontheroad?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Gabriella Clare Marino</a> on <a href="https://unsplash.com/?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText">Unsplash</a>
  