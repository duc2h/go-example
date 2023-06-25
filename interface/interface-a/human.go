package interfacea

type Human interface {
	Talk(s string)
	Eat(food string)
}

func CallTalk(h Human) {
	h.Talk("CallTalk ")
}

func CallEat(h Human) {
	h.Eat("CallEat ")
}

type RunHuman struct {
	human []Human
}

func (r *RunHuman) WithHuman(h ...Human) {
	r.human = append(r.human, h...)
}

func (r RunHuman) RunTalk(s string) {
	for _, h := range r.human {
		h.Talk(s)
	}
}

func (r RunHuman) RunFood(food string) {
	for _, h := range r.human {
		h.Eat(food)
	}
}
