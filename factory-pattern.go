package main

import "fmt"

type Product interface {
	operation()
}

type ConcreteProductA struct{}

func (a *ConcreteProductA) operation() {
	fmt.Println("Concrete Product A")
}

type ConcreteProductB struct{}

func (b *ConcreteProductB) operation() {
	fmt.Println("Concrete Product B")
}

type Creator interface {
	factoryMethod() Product
}

type CreatorA struct{}

func (a *CreatorA) factoryMethod() Product {
	return &ConcreteProductA{}
}

type CreatorB struct{}

func (b *CreatorB) factoryMethod() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &CreatorA{}
	productA := creatorA.factoryMethod()
	productA.operation()

	creatorB := &CreatorB{}
	productB := creatorB.factoryMethod()
	productB.operation()
}
