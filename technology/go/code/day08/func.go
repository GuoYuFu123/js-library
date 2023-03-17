package main

import (
	"fmt"

	"../utils"
)

func main() {
	fmt.Println("model")

	var stu = utils.NewStudent("xiaoming", 12)
	fmt.Println("stu", *stu, "age=", stu.GetAge())

}
