package main

import (
	"fmt"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	fmt.Println("stat")
	isExist, err := PathExist("./w1.txt")
	fmt.Println("isExist", isExist)
	fmt.Println("err", err)

}
