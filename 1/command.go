package main

import "fmt"

type Command interface {
	Execute()
}

type Gate struct {
}

func (g Gate)Execute() {
	fmt.Println("The gate opens")
}

type Ventilator struct {
}

func (g Ventilator)Execute() {
	fmt.Println("The ventilator running")
}

type Terminal struct {
	ventilator Command
	gate Command
}

func NewTerminal(ventilator Command, gate Command) *Terminal {
	return &Terminal{
		ventilator: ventilator,
		gate: gate,
	}
}

func (t *Terminal) RunVentilation() {
	t.ventilator.Execute()
}

func (t *Terminal) OpenGate() {
	t.gate.Execute()
}