package main

import "fmt"

type Pills interface {
	TakePills()
}

type ibuprofen struct {
}

func (i ibuprofen)TakePills() {
	fmt.Println("Taking Ibuprofen")
}

type novoPassit struct {
}

func (i novoPassit)TakePills() {
	fmt.Println("Taking Novo-Passit")
}

type ascorbicAcid struct {
}

func (i ascorbicAcid)TakePills() {
	fmt.Println("Taking Ascorbic Acid")
}

func PillsFactory(name string) Pills {
	switch name {
	case "ибупрофен":
		return ibuprofen{}
	case "ново-пассит":
		return novoPassit{}
	default:
		return ascorbicAcid{}
	}
}