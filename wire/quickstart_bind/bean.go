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

//===================App

type App struct {
	Event      Event
	AppService AppService
}

func NewApp(e Event, svr AppService) App {
	return App{e, svr}
}

func (app App) Start() {
	greetMsg := app.Event.Greeter.Greet()
	installMsg := app.AppService.Install()
	fmt.Println(greetMsg)
	fmt.Println(installMsg)
}
