package main

import (
	"fmt"
)

type vehicles interface {
}
type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	p := car{}
	t := car{}
	b := car{}
	cars := []car{p, t, b}

	b747 := plane{}
	b757 := plane{}
	b767 := plane{}

	planes := []plane{b747, b757, b767}

	s := boat{}
	n := boat{}
	m := boat{}
	boats := []boat{s, n, m}

	for k, v := range cars {
		fmt.Println(k, " - ", v)
	}
	for k, v := range planes {
		fmt.Println(k, " - ", v)
	}
	for k, v := range boats {
		fmt.Println(k, " - ", v)
	}

	fmt.Println("---------------")
	rides := []vehicles{p, t, b, b747, b757, b767, s, n, m}
	for k, v := range rides {
		fmt.Println(k, " - ", v)
	}
}
