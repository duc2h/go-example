package main

import (
	impla "github-com/edarha/go-example/interface/implement-a"
	implb "github-com/edarha/go-example/interface/implement-b"
	inter "github-com/edarha/go-example/interface/interface-a"
)

func main() {

	// option 1
	per1 := impla.NewPerson()
	per2 := implb.NewPerson()

	human := inter.RunHuman{}
	human.WithHuman(per1, per2)

	human.RunTalk("hey!")
	human.RunFood("delicious")

	// option 2
	inter.CallTalk(per1)
	inter.CallTalk(per2)
	inter.CallEat(per1)
	inter.CallEat(per2)
}
