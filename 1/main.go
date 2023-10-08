package main

import (
	"fmt"
)

type MyStruct struct {
	MyField      int ""
	MyOtherField int `what:"isThis"`
}

func main() {
	// Facade:
	fmt.Println("#% Facade pattern: ")
	talkFacade := NewStoreFacade()
	talkFacade.SaveToStore("Some data")
	fmt.Println(" <------------> ")

	// Builder:
	fmt.Println("#% Builder pattern: ")
	b := NewUserBuilder()
	b.EnableCasher().SetCasher("Logger").SetStoreToSave("Postgres").Build()
	fmt.Println(" <------------> ")

	// Visitor:
	fmt.Println("#% Visitor pattern: ")
	visitor := NewWorker()
	visitor.DoSomeWork()
	visitor.append(WorkChangerToHard{})
	visitor.DoSomeWork()
	visitor.append(WorkChangerToEaasier{})
	visitor.DoSomeWork()
	fmt.Println(" <------------> ")

	// Command:
	fmt.Println("#% Command pattern: ")
	ventilatorCommand := Ventilator{}
	gateCommand := Gate{}
	terminal := NewTerminal(ventilatorCommand, gateCommand)
	terminal.OpenGate()
	terminal.RunVentilation()
	fmt.Println(" <------------> ")

	// Chain-of-responsibility:
	fmt.Println("#% Chain-of-responsibility: ")
	compbuilder := NewComputerBuilder()
	comp := NewComputer()
	compbuilder.execute(comp)
	fmt.Println(" <------------> ")

	// Factory method:
	fmt.Println("#% Factory method: ")
	pills := PillsFactory("ново-пассит")
	pills.TakePills()
	pills = PillsFactory("ибупрофен")
	pills.TakePills()
	pills = PillsFactory("незнамо что")
	pills.TakePills()
	fmt.Println(" <------------> ")

	// Strategy:
	fmt.Println("#% Strategy: ")
	quick := QuickSort{}
	st := NewStore(quick)
	st.Add(10)
	merge := MergerSort{}
	st.SetSort(merge)
	st.Add(12)
	fmt.Println(" <------------> ")
}
