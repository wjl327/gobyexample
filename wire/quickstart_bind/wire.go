// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

// wire.go

var EventProvider = wire.NewSet(NewEvent, NewMessage, NewGreeter)
var ServiceProvider = wire.NewSet(NewCoreAppService, wire.Bind(new(AppService), new(*CoreAppService)))

func InitializeApp() App {
	wire.Build(NewApp, EventProvider, ServiceProvider)
	return App{}
}
