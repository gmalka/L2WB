package main

import "fmt"

type WorkChangerToHard struct {
}

func (wc WorkChangerToHard) WorkHarder(w *Worker) {
	w.workKind = "HARD"
}

func (wc WorkChangerToHard) WorkEasier(w *Worker) {
}

type WorkChangerToEaasier struct {
}

func (wc WorkChangerToEaasier) WorkHarder(w *Worker) {
}

func (wc WorkChangerToEaasier) WorkEasier(w *Worker) {
	w.workKind = "EASY"
}

type Visitor interface {
	WorkHarder(*Worker)
	WorkEasier(*Worker)
}

func NewWorker() *Worker {
	return &Worker{}
}

type Worker struct {
	workKind string
}

func (v *Worker) append(visitor Visitor) {
	visitor.WorkHarder(v)
	visitor.WorkEasier(v)
}

func (v *Worker) DoSomeWork() {
	fmt.Printf("Let's do some %s work\n", v.workKind)
}
