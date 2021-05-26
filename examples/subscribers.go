package main

import (
	"fmt"
)

type (
	subscriberA struct {
	}
	subscriberB struct {
	}
	subscriberC struct {
	}
)

func (s *subscriberA) receiveChanA() {
	fmt.Println("Subscriber A receiving")
}

func (s *subscriberB) receiveChanB(data string) {
	fmt.Println("Subscriber B receiving", data)
}

func (s *subscriberC) receiveChanC(data Person) {
	fmt.Println("Subscriber C receiving", data)
}
