package event

import (
	"errors"
	"fmt"
)

type Message string

func NewMessage(phrase string) (Message, error) {
	if phrase == "" {
		return Message(""), errors.New("phrase was not provided")
	}

	return Message("Hello, World!"), nil
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) (*Greeter, error) {
	if m == "" {
		return nil, errors.New("message was not provided")
	}
	return &Greeter{Message: m}, nil
}

func (g *Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter *Greeter
}

func NewEvent(g *Greeter) (*Event, error) {
	if g == nil {
		return nil, errors.New("greeter was not provided")
	}
	return &Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
