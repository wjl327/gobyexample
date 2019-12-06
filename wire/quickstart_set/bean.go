package main

import "fmt"

//===================Message

type Message string

func NewMessage() Message {
	return Message("Hello!")
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

//===================Ruder

type Ruder string

func NewRuder() Ruder {
	return Ruder("Go away!")
}

//===================App

type App struct {
	Event Event
	Ruder Ruder
}

func NewApp(e Event, r Ruder) App {
	return App{e, r}
}

func (app App) Start() {
	greetMsg := app.Event.Greeter.Greet()
	ruderMsg := app.Ruder
	fmt.Println(greetMsg)
	fmt.Println(ruderMsg)
}
