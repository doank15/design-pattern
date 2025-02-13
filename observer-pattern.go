package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

type ConcreteObserverA struct {
	name string
}

func (a *ConcreteObserverA) Update(message string) {
	fmt.Printf("%s received message: %s\n\n", a.name, message)
}

func main() {
	subject := Subject{}
	observerA := ConcreteObserverA{name: "Observer A"}

	subject.Attach(&observerA)
	subject.Notify("State A")
}
