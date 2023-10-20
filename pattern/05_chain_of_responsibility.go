package main

import "fmt"

type chain interface {
	Executer
	setNext(chain) chain
}

type Executer interface {
	execute(*Computer)
}

type computerCaseInstaller struct {
	next chain
}

func (c *computerCaseInstaller) execute(comp *Computer) {
	if comp.compcase {
		fmt.Println("Computer case already installed")
		c.next.execute(comp)
	} else {
		fmt.Println("Installing a Computer case")
		c.next.execute(comp)
	}

}

func (c *computerCaseInstaller) setNext(next chain) chain {
	c.next = next
	return next
}

type motherboardInstaller struct {
	next chain
}

func (c *motherboardInstaller) execute(comp *Computer) {
	if comp.motherboard {
		fmt.Println("Motherboard already installed")
		c.next.execute(comp)
	} else {
		fmt.Println("Installing a Motherboard")
		c.next.execute(comp)
	}

}

func (c *motherboardInstaller) setNext(next chain) chain {
	c.next = next
	return next
}

type cpuInstaller struct {
	next chain
}

func (c *cpuInstaller) execute(comp *Computer) {
	if comp.cpu {
		fmt.Println("CPU already installed")
		c.next.execute(comp)
	} else {
		fmt.Println("Installing a CPU")
		c.next.execute(comp)
	}

}

func (c *cpuInstaller) setNext(next chain) chain {
	c.next = next
	return next
}

type videocardInstaller struct {
	next chain
}

func (c *videocardInstaller) execute(comp *Computer) {
	if comp.videocard {
		fmt.Println("Videocard already installed")
	} else {
		fmt.Println("Installing a Videocard")
	}
	fmt.Println("Computer building finish, enjoy!")
}

func (c *videocardInstaller) setNext(next chain) chain {
	c.next = next
	return next
}

func NewComputerBuilder() Executer {
	var cs, origin chain
	origin = &computerCaseInstaller{}
	cs = origin.setNext(&motherboardInstaller{})
	cs = cs.setNext(&cpuInstaller{})
	_ = cs.setNext(&videocardInstaller{})

	return origin
}

type Computer struct {
	compcase    bool
	motherboard bool
	cpu         bool
	videocard   bool
}

func NewComputer() *Computer {
	return &Computer{}
}