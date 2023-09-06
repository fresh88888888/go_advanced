package main

import (
	"fmt"
	"reflect"
)

type Handler func() *int
type VarHandler func(...int)
type intHandler func(int) int

func compose(a intHandler, b intHandler) intHandler {
	return func(c int) int {
		return a(b(c))
	}
}

func init() {
	fmt.Println("Call-1")
}

func init() {
	fmt.Println("Call-2")
}

func main() {

	var f Handler = func() *int {
		i := 1
		return &i
	}

	fmt.Println(&f)

	fmt.Println(reflect.TypeOf(f).Comparable()) //false

	var vh VarHandler = func(i ...int) {
		fmt.Println(i)
	}

	vh([]int{1, 2, 3}...)

	add2 := compose(func(i int) int {
		return i + 1
	}, func(i int) int {
		return i + 1
	})

	fmt.Println(add2(0)) // 2
}
