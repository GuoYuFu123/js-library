package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("reflect change val")

	var num = 100;
	fn := reflect.ValueOf(&num)
	fmt.Println("fn", fn)
	fn.Elem().SetInt(200)
	fmt.Println("num=", num)
}