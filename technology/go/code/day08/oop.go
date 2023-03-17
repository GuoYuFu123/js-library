package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) eat() {
	fmt.Println(animal.Name, ":eat...")
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

type A struct {
	Name string
}

type B struct {
	Name string
}
type C struct {
	A
	B
}
type D struct {
	a A
	B
}

type E struct {
	Name string
	int
}

type G struct {
	C
	Name string
}

func main() {
	fmt.Println("oop")

	cat := Cat{}
	cat.Name = "cat"
	cat.Animal.Age = 10
	cat.eat()

	dog := Dog{}
	dog.Name = "dog"
	dog.eat()
	dog.Animal.eat()

	c := C{}
	c.A.Name = "cc"
	d := D{}
	d.a.Name = "dd"

	e := E{"张三", 100}
	fmt.Println("e=", e)

	g := G{}
	g.Name = "g-name"
	g.C.A.Name = "gca-name"
	fmt.Println("g=", g)

	g2 := G{Name: "name", C: C{A: A{Name: "aa"}, B: B{Name: "bb"}}}
	fmt.Println("g2=", g2.C.A.Name)
}
