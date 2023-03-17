package Service

import (
	"CustomerManage/Model"
	"fmt"
)

// 声明一个CustomerServer结构体，该结构体完成对Customer的CRUD操作
type CustomerService struct {
	customer    []Model.Customer //用于存储多个Customer对象的切片
	customerNum int              //用于统计存储Customer对象个数的属性
}

// 使用一个工厂函数，用于返回一个CustomerService对象
func NewCustomerService() *CustomerService {
	//初始化一个客户信息
	service := CustomerService{
		customerNum: 1,
	}
	customer := Model.NewCustomer(1, "张三", "男", 30, "40066668888", "zhangsan@qq.com") //工厂函数创建一个Customer
	service.customer = append(service.customer, customer)                             //追加到切片
	return &service
}

// 返回客户信息的切片
func (this *CustomerService) GetCustomers() []Model.Customer {
	return this.customer //所有的customer对象，切片类型
}

// 添加客户
func (this *CustomerService) Add(name string, sex string, age int,
	phone string, email string) bool {
	this.customerNum++                                                            //id动态增长
	customer := Model.NewCustomer(this.customerNum, name, sex, age, phone, email) //工厂函数创建一个Customer
	this.customer = append(this.customer, customer)                               //追加到切片
	return true
}

// 删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindIndexById(id)
	if index != -1 {
		this.customer = append(this.customer[:index], this.customer[index+1:]...) //去掉下标为index的切片
		return true                                                               //删除成功
	}
	return false //删除失败
}

// 删除客户
func (this *CustomerService) Update(id int) bool {
	//查看该用户是否存在
	if this.FindIndexById(id) == -1 {
		fmt.Println("该用户不存在")
		return false
	}

	customer := this.FindCustomerById(id) //获取该用户信息
	var name string
	var sex string
	var age int
	var phone string
	var email string
	fmt.Printf("姓名(%v):", customer.Name)
	fmt.Scanln(&name)
	if name != "" {
		customer.Name = name
	}
	fmt.Printf("性别(%v):", customer.Gender)
	fmt.Scanln(&sex)
	if sex != "" {
		customer.Gender = sex
	}
	fmt.Printf("年龄(%v):", customer.Age)
	fmt.Scanln(&age)
	if age != 0 {
		customer.Age = age
	}
	fmt.Printf("电话(%v):", customer.Phone)
	fmt.Scanln(&phone)
	if phone != "" {
		customer.Phone = phone
	}
	fmt.Printf("邮箱(%v):", customer.Email)
	fmt.Scanln(&email)
	if email != "" {
		customer.Email = email
	}
	return true
}

// 通过id查找客户在切片当中存储的下标，未找到返回-1
func (this *CustomerService) FindIndexById(id int) int {
	for index, value := range this.customer {
		if value.Id == id {
			return index //找到
		}
	}
	return -1 //未找到
}

// 通过id查找客户信息
func (this *CustomerService) FindCustomerById(id int) *Model.Customer {
	for i := 0; i < len(this.customer); i++ {
		if this.customer[i].Id == id {
			return &this.customer[i]
		}
	}
	return &Model.Customer{} //未找到
}
