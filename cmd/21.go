package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct {
	message string
}

func NewAdaptee(message string) *Adaptee {
	return &Adaptee{message}
}

func (a *Adaptee) SpecificRequest() string {
	return a.message
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := NewAdaptee("Это адаптируемая структура")
	adapter := NewAdapter(adaptee)
	result := adapter.Request()
	fmt.Println(result)
}
