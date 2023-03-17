package Model

import "fmt"

// 声明一个Customer结构体，表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 使用一个工厂函数，用于返回一个Customer对象
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 重写String方法
func (this *Customer) String() string {
	var info string = fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
		this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
