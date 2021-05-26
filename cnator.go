package cnator

import (
	"reflect"
)

type (
	Cnator struct {
		subscriptions map[interface{}][]interface{}
	}
)

func New() *Cnator {
	return &Cnator{
		subscriptions: make(map[interface{}][]interface{}),
	}
}

func (c *Cnator) Subscribe(channel interface{}, subscriber interface{}) {
	channelType := validateChannel(channel)
	validateSubsciber(subscriber, channelType)

	c.subscriptions[channel] = append(c.subscriptions[channel], subscriber)
}

func (c *Cnator) Serve() {
	for channel, subscibers := range c.subscriptions {

		go func(channel interface{}, subscibers []interface{}) {
			channelVal := reflect.ValueOf(channel)

			for {
				channelData, ok := channelVal.Recv()
				if !ok {
					return
				}

				for _, subscriber := range subscibers {
					subcriberVal := reflect.ValueOf(subscriber)

					channelType := reflect.TypeOf(channel)

					args := []reflect.Value{}

					if channelType.Elem().Name() != "" {
						args = []reflect.Value{channelData}
					}

					subcriberVal.Call(args)

				}
			}
		}(channel, subscibers)
	}
}

func validateChannel(channel interface{}) reflect.Type {
	channelType := reflect.TypeOf(channel)

	if channelType.Kind() != reflect.Chan {
		panic("channel parameter is not a channel")
	}

	if reflect.ValueOf(channel).IsNil() {
		panic("channel is uninitialized")
	}
	return channelType
}

func validateSubsciber(subscriber interface{}, channelType reflect.Type) {
	if reflect.ValueOf(subscriber).IsNil() {
		panic("subscriber is uninitialized")
	}

	subscriberType := reflect.TypeOf(subscriber)

	if subscriberType.Kind() != reflect.Func {
		panic("subscriber must be a function")
	}

	channelParamType := channelType.Elem().Name()
	subscriberParams := subscriberType.NumIn()

	if channelParamType == "" {
		if subscriberParams != 0 {
			panic("channel is empty struct, so subscriber should not have any parameter")
		}
	} else {
		if subscriberParams != 1 {
			panic("subscriber must have one and only one parameter")
		}

		subscriberParamType := subscriberType.In(0).Name()
		if channelParamType != subscriberParamType {
			panic("subscriber " + subscriberType.String() + " has different param type than the channel: " + channelParamType)
		}
	}
}
