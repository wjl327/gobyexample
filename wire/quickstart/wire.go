// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

// wire.go

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
