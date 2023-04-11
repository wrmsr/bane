package unions

import (
	"fmt"
	"testing"
)

//

type Howler interface {
	Howl()
}

type LoudHowl struct{}

func (h LoudHowl) Howl() { fmt.Println("HOWL") }

type QuietHowl struct{}

func (h QuietHowl) Howl() { fmt.Println("howl") }

type Dog struct {
	Woofs int
	Howl  Howler
}

type Person struct {
	Age   float32
	Money int
}

//

type Union1_Kind int8

const (
	Union1_Kind_Dog Union1_Kind = iota + 1
	Union1_Kind_Person
)

func (k Union1_Kind) String() string {
	switch k {
	case Union1_Kind_Dog:
		return "Dog"
	case Union1_Kind_Person:
		return "Person"
	}
	return "<invalid>"
}

type Union1 struct {
	k Union1_Kind

	int_0     int
	float32_0 float32
	any_0     any
}

func (u Union1) Kind() Union1_Kind { return u.k }

func Union1_Dog(v Dog) Union1 {
	u := Union1{k: Union1_Kind_Dog}
	u.int_0 = v.Woofs
	if v.Howl != nil {
		u.any_0 = v.Howl
	}
	return u
}

func (u Union1) Dog() Dog {
	if u.k == Union1_Kind_Dog {
		return Dog{
			Woofs: u.int_0,
			Howl:  u.any_0.(Howler),
		}
	}
	panic(u.k)
}

func Union1_Person(v Person) Union1 {
	u := Union1{k: Union1_Kind_Person}
	u.float32_0 = v.Age
	u.int_0 = v.Money
	return u
}

func (u Union1) Person() Person {
	if u.k == Union1_Kind_Person {
		return Person{
			Age:   u.float32_0,
			Money: u.int_0,
		}
	}
	panic(u.k)
}

func TestUnion1(t *testing.T) {
	var u Union1

	u = Union1_Dog(Dog{Woofs: 5})
	fmt.Printf("%#v\n", u.Dog())

	u = Union1_Dog(Dog{Woofs: 5, Howl: LoudHowl{}})
	fmt.Printf("%#v\n", u.Dog())

	u = Union1_Person(Person{Age: 23.4, Money: 420})
	fmt.Printf("%#v\n", u.Person())
}
