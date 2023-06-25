package implementa

import "fmt"

type person1 struct {
}

func NewPerson() *person1 {
	return &person1{}
}

func (p *person1) Talk(s string) {
	fmt.Println("person1 is talking: ", s)
}

func (p *person1) Eat(food string) {
	fmt.Println("person1 is eating: ", food)
}
