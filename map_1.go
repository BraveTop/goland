package main

import (
	"fmt"
)

func a() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("it is only one")
		}

	}()
	panic("panic a")
}

func main() {
	a()
	fmt.Println("hello world")
}
