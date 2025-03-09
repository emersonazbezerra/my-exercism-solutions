package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(a string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", i.LanguageName(), a)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(a string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", p.LanguageName(), a)
}
