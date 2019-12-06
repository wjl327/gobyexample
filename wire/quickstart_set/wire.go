// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

// wire.go

var EventSet = wire.NewSet(NewEvent, NewMessage, NewGreeter)
var AppSet = wire.NewSet(NewRuder)

func InitializeApp() App {
	wire.Build(NewApp, AppSet, EventSet)
	return App{}
}
