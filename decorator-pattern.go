package main

import (
	"fmt"
	"time"
)

// Service to create a service
type Service interface {
	FetchData(url string) string
}

// ConcreteServiceA to implement service
type ConcreteServiceA struct {
}

func (s *ConcreteServiceA) FetchData(url string) string {
	return fmt.Sprintf("Data for: %s", url)
}

// LoggingDecorator to create a new logging service without modifying object ConcreteServiceA
type LoggingDecorator struct {
	WrappedObject Service
}

// FetchData to implement service
func (l LoggingDecorator) FetchData(param string) string {
	start := time.Now()

	data := l.WrappedObject.FetchData(param)

	fmt.Printf("Fetched data from %s to %s\n", time.Since(start))

	return data
}

func main() {
	service := ConcreteServiceA{}
	// Wrap an object with multiple decorators
	logging := LoggingDecorator{WrappedObject: &service}

	result := logging.WrappedObject.FetchData("123")

	fmt.Println(result)
}
