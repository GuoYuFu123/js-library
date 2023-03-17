package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

func (student Student) String() string {
	return fmt.Sprintf("guoguo:Name=[%v] Age=[%v]", student.Name, student.Age)
}

func main() {
	fmt.Println("int")

	var i integer = 1
	i.print()
	i.change()
	fmt.Println("i==", i)

	stu := Student{"xiaoming", 12}
	fmt.Println("stu", &stu)
}
