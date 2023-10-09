package main

import "fmt"

type State interface {
	DoHardWork()
	DoRegularWork()
}

func NewMechanism() *mechanism {
	mechanism := &mechanism{}
	regularState := &regularWorkedState{
		m: mechanism,
	}

	safetyState := &safetyWorkedState{
		m: mechanism,
	}

	mechanism.regular = regularState
	mechanism.safety = safetyState
	mechanism.current = regularState

	return mechanism
}

type mechanism struct {
	regular State
	safety State

	current State

	temperature int
}

func (m *mechanism)SetState(s State) {
	m.current = s
}

func (m *mechanism)Wait() {
	if m.temperature > 0 {
		m.temperature -= 40
		if m.temperature < 0 {
			m.temperature = 0
		}

		if m.temperature < 50 {
			fmt.Println("Switching to regular state")
			m.SetState(m.regular)
		}
		fmt.Printf("Mechanism waiting, current temperature: %d\n", m.temperature)
	} else {
		fmt.Println("Mechanisms temperature already are 0")
	}
}

func (m *mechanism)DoHardWork() {
	m.current.DoHardWork()
}

func (m *mechanism)DoRegularWork() {
	m.current.DoRegularWork()
}

type regularWorkedState struct {
	m *mechanism
}

func (h *regularWorkedState)DoHardWork() {
	if h.m.temperature + 30 >= 100 {
		fmt.Println("Switch to Safety state")
		h.m.SetState(h.m.safety)
		h.m.DoHardWork()
	} else {
		fmt.Println("Making hard work by regular state")
		h.m.temperature += 30
	}
}

func (h *regularWorkedState)DoRegularWork() {
	if h.m.temperature + 15 >= 100 {
		fmt.Println("Switch to Safety state")
		h.m.SetState(h.m.safety)
		h.m.DoRegularWork()
	} else {
		fmt.Println("Making regular work by regular state")
		h.m.temperature += 15
	}
}

type safetyWorkedState struct {
	m *mechanism
}

func (h *safetyWorkedState)DoHardWork() {
	if h.m.temperature + 10 > 100 {
		fmt.Println("Too hot condition to hard work, please wait")
	} else {
		fmt.Println("Making hard work by safety state")
		h.m.temperature += 10
	}
}

func (h *safetyWorkedState)DoRegularWork() {
	if h.m.temperature + 5 > 100 {
		fmt.Println("Too hot condition to regular work, please wait")
	} else {
		fmt.Println("Making regular work by safety state")
		h.m.temperature += 5
	}
}