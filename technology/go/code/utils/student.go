package utils

type student struct {
	Name string
	age  float64
}

func NewStudent(name string, age float64) *student {
	return &student{
		Name: name,
		age:  age,
	}
}

func (s *student)GetAge() float64{
	return s.age 
}