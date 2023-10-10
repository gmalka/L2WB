package main

import "fmt"

// Фасад, который инкапсулирует низкоуровневые элементы
type StoreFacade struct {
	h *Cashe
	s *Memory
}

func NewStoreFacade() StoreFacade {
	return StoreFacade{
		h: NewCashe(),
		s: NewMemory(),
	}
}

func (t StoreFacade)SaveToStore(str string) {
	t.h.SaveToCashe(str)
	t.s.SaveToMemory(str)
}

// Низкоуровневые элементы
type Cashe struct{}
type Memory struct{}

func NewCashe() *Cashe {
	h := Cashe(struct{}{})
	return &h
}

func NewMemory() *Memory {
	s := Memory(struct{}{})
	return &s
}

func (h *Cashe)SaveToCashe(str string) {
	fmt.Printf("%s saved to Cashe\n", str)
}

func (s *Memory)SaveToMemory(str string) {
	fmt.Printf("%s saved to Memory\n", str)
}