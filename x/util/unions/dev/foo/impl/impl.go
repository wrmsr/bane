//go:build !nodev

package impl

///

type Howler interface {
	Howl() string
}

type LoudHowl struct{}

func (h LoudHowl) Howl() string { return "HOWL" }

type QuietHowl struct{}

func (h QuietHowl) Howl() string { return "howl" }

type Dog struct {
	Woofs int
	Howl  Howler
}

type Person struct {
	Age   float32
	Money int
	Name  string
}

///

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
	string_0  string
}

func (u Union1) Kind() Union1_Kind { return u.k }

//

func Union1_Dog(v Dog) Union1 {
	u := Union1{k: Union1_Kind_Dog}
	u.int_0 = v.Woofs
	if v.Howl != nil {
		u.any_0 = v.Howl
	}
	return u
}

func (u Union1) Dog() Dog {
	if u.k != Union1_Kind_Dog {
		panic(u.k)
	}
	v := Dog{}
	v.Woofs = u.int_0
	if u.any_0 != nil {
		v.Howl = u.any_0.(Howler)
	}
	return v
}

func (u Union1) Dog_Woofs() int {
	if u.k != Union1_Kind_Dog {
		panic(u.k)
	}
	return u.int_0
}

func (u Union1) Dog_Howl() Howler {
	if u.k != Union1_Kind_Dog {
		panic(u.k)
	}
	if u.any_0 != nil {
		return u.any_0.(Howler)
	}
	return nil
}

type Union1_View_Dog struct {
	u *Union1
}

func (uv Union1_View_Dog) Woofs() int {
	return uv.u.int_0
}

func (uv Union1_View_Dog) Howl() Howler {
	if uv.u.any_0 != nil {
		return uv.u.any_0.(Howler)
	}
	return nil
}

func (u Union1) View_Dog() Union1_View_Dog {
	if u.k != Union1_Kind_Dog {
		panic(u.k)
	}
	return Union1_View_Dog{u: &u}
}

//

func Union1_Person(v Person) Union1 {
	u := Union1{k: Union1_Kind_Person}
	u.float32_0 = v.Age
	u.int_0 = v.Money
	u.string_0 = v.Name
	return u
}

func (u Union1) Person() Person {
	if u.k != Union1_Kind_Person {
		panic(u.k)
	}
	v := Person{}
	v.Age = u.float32_0
	v.Money = u.int_0
	v.Name = u.string_0
	return v
}

///
