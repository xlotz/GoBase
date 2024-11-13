package main

import (
	"log"
	"testing"
	"testing/quick"
)

func TestAdder(t *testing.T) {
	res := Adder(4, 7)
	if res != 11 {
		t.Fatal("not 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

func ExampleAdder() {
	log.Println(Adder(4, 7))
}

func ExampleAdder_multiple() {
	log.Println(Adder(3, 6, 7, 4, 61))
}

func TestAdderBlackbox(t *testing.T) {
	err := quick.Check(a, nil)

	if err != nil {
		t.Fatal(err)
	}
}
func a(x, y int) bool {
	return Adder(x, y) == x+y
}
