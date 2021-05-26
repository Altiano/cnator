package main

import "strconv"

type (
	producerA struct {
		chanA chan<- struct{}
	}

	producerB struct {
		chanB   chan<- string
		counter int
	}

	producerC struct {
		chanC   chan<- Person
		counter int
	}
)

func NewProducerA(chanA chan<- struct{}) *producerA {
	return &producerA{
		chanA: chanA,
	}
}

func (p *producerA) publish() {
	p.chanA <- struct{}{}
}

func NewProducerB(chanB chan<- string) *producerB {
	return &producerB{
		chanB: chanB,
	}
}

func (p *producerB) publish() {
	p.counter++
	p.chanB <- "string #" + strconv.Itoa(p.counter)
}

func NewProducerC(chanC chan<- Person) *producerC {
	return &producerC{
		chanC: chanC,
	}
}

func (p *producerC) publish() {
	p.counter++
	p.chanC <- Person{
		name:  "johndoe" + strconv.Itoa(p.counter),
		email: "johndoe" + strconv.Itoa(p.counter) + "@email.com",
	}
}
