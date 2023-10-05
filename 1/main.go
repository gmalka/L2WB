package main

import "fmt"

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
}