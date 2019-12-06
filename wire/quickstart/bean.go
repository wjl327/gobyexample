package main

import "fmt"

//===================Message

type Message string

func NewMessage() Message {
	return Message("Hello")
}

//===================Greeter

type Greeter struct {
	Message Message // <- adding a Message field
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

//===================Event

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
