package main

import "fmt"

type Usb interface {
	start()
	stop()
}
type Phone struct{}

func (p Phone) start() {
	fmt.Println("phone start...")
}
func (p Phone) stop() {
	fmt.Println("phone stop...")
}

func (p Phone) call() {
	fmt.Println("phone call···")
}

type Camera struct{}

func (c Camera) start() {
	fmt.Println("camrea start...")
}
func (c Camera) stop() {
	fmt.Println("camrea stop...")
}

type Computer struct{}

func (c Computer) Working(usb Usb) {
	usb.start()
	// 类型断言的实现，注意体会
	if phone, ok := usb.(Phone); ok {
		phone.call();
	}
	usb.stop()
}

func main() {
	fmt.Println("interface")
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

	var b Usb = phone
	b.start()

}
