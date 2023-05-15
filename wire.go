//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ofekatr/go-wire-practice/event"
)

func InitializeEvent(phrase string) (*event.Event, error) {
	wire.Build(event.NewEvent, event.NewGreeter, event.NewMessage)
	return &event.Event{}, nil
}
