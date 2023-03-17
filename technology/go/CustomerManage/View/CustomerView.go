package View

import (
	"CustomerManage/Service"
	"fmt"
)

// 声明CostomerView结构体
type ConstomerView struct {
	Key             string                   //用于接收用户输入的字段
	Loop            bool                     //表示是否循环的显示主菜单
	CustomerService *Service.CustomerService //用于存储CustomerService对象的属性
}

// 用于床架CustomerView的工厂函数
func NewCustomerView() *ConstomerView {
	//创建CustomerView
	customerView := ConstomerView{
		Key:             "",
		Loop:            true,
		CustomerService: Service.NewCustomerService(),
	}
	return &customerView
}

// 用于显示主菜单的方法
func (this *ConstomerView) MainMenu() {
	for {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Println("\t\t1 添 加 客 户")
		fmt.Println("\t\t2 修 改 客 户")
		fmt.Println("\t\t3 删 除 客 户")
		fmt.Println("\t\t4 客 户 列 表")
		fmt.Println("\t\t5 退       出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&this.Key)
		switch this.Key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("您输入的选择有误，请重新输入！")
		}
		if !this.Loop {
			break
		}
	}
}

// 显示客户列表
func (this *ConstomerView) list() {
	//获取所有客户的信息
	customers := this.CustomerService.GetCustomers()
	//显示
	fmt.Println("---------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, value := range customers {
		fmt.Println(&value) //Customer对象重写过String()方法，默认调用String()方法
	}
	fmt.Println("\n-----------客户列表显示完成-----------\n\n")
}

// 添加客户
func (this *ConstomerView) add() {
	var name string
	var sex string
	var age int
	var phone string
	var email string
	fmt.Println("---------------添加客户---------------")

	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("性别：")
	fmt.Scanln(&sex)

	fmt.Print("年龄：")
	fmt.Scanln(&age)

	fmt.Print("电话：")
	fmt.Scanln(&phone)

	fmt.Print("邮箱：")
	fmt.Scanln(&email)
	if this.CustomerService.Add(name, sex, age, phone, email) {
		fmt.Println("---------------添加完成---------------")
	} else {
		fmt.Println("---------------添加失败---------------")
	}
}

// 删除客户
func (this *ConstomerView) delete() {
	fmt.Println("---------------删除客户---------------")
	var id int
	fmt.Print("请选择待删除的客户编号(-1退出)：")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("退出了删除客户！~")
		return
	}
	var key string
	fmt.Print("确认是否删除(Y/N)：")
	fmt.Scanln(&key)
	if key == "Y" || key == "y" {
		if this.CustomerService.Delete(id) {
			fmt.Println("---------------删除完成---------------")
		} else {
			fmt.Println("---------------删除失败,未找到该客户信息---------------")
		}
	} else {
		fmt.Println("---------------取消删除---------------")
	}
}

// 修改客户信息
func (this *ConstomerView) update() {
	fmt.Println("---------------修改客户---------------")
	var id int
	fmt.Print("请选择待修改的客户编号(-1退出)：")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("退出了修改客户！~")
		return
	}

	if this.CustomerService.Update(id) {
		fmt.Println("---------------修改完成---------------")
	} else {
		fmt.Println("---------------修改失败---------------")
	}
}

func (this *ConstomerView) exit() {
	var key string
	for {
		fmt.Print("确认是否退出(Y/N)：")
		fmt.Scanln(&key)
		if key == "Y" || key == "y" {
			fmt.Println("退出成功！~")
			this.Loop = false
			return
		} else if key == "N" || key == "n" {
			return
		}
	}
}
