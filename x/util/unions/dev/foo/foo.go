//go:build !nodev

package main

import fi "github.com/wrmsr/bane/x/util/unions/dev/foo/impl"

//go:noinline
func frob(u fi.Union1) string {
	switch u.Kind() {
	case fi.Union1_Kind_Dog:
		return u.Dog().Howl.Howl()
	case fi.Union1_Kind_Person:
		return u.Person().Name
	default:
		panic(u)
	}
}

//go:noinline
func frob0(u fi.Union1) string {
	return u.Dog().Howl.Howl()
}

//go:noinline
func frob1(u fi.Union1) string {
	return u.Dog_Howl().Howl()
}

//go:noinline
func frob2(u fi.Union1) string {
	return u.View_Dog().Howl().Howl()
}

func main() {
	frob(fi.Union1_Dog(fi.Dog{Woofs: 5, Howl: fi.LoudHowl{}}))
	frob(fi.Union1_Dog(fi.Dog{Woofs: 5, Howl: fi.QuietHowl{}}))
	frob(fi.Union1_Person(fi.Person{Age: 23.4, Money: 420}))

	d := fi.Union1_Dog(fi.Dog{Woofs: 5, Howl: fi.LoudHowl{}})
	frob0(d)
	frob1(d)
	frob2(d)
}
