package implementb

import "fmt"

type person2 struct {
}

func NewPerson() *person2 {
	return &person2{}
}

func (p *person2) Talk(s string) {
	fmt.Println("person2 is talking: ", s)
}

func (p *person2) Eat(food string) {
	fmt.Println("person2 is eating: ", food)
}
