package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type intPtr *int
type chanType chan bool

type rightChan chan<- bool
type leftChan chan<- bool

type uPoint unsafe.Pointer

func main() {
	n := 1
	var pn intPtr = &n
	var up = uPoint(&pn)

	fmt.Println(reflect.TypeOf(pn).Comparable()) // true
	fmt.Println(reflect.TypeOf(pn))              // intPtr
	fmt.Println(reflect.TypeOf(up))

	fmt.Println(reflect.TypeOf(nil)) // nil

	// Underlying type
	type IntSlice []int
	type customSlice IntSlice
	type anotherSlice = customSlice
	type anotehrCustomSlice customSlice

	fmt.Println(reflect.TypeOf([]customSlice{}).Comparable()) //false
	fmt.Println(reflect.TypeOf([]anotherSlice{}))             // []main.customerSlice

	var is = IntSlice{}
	var cs = customSlice{}
	var as = anotherSlice{}
	var acs = anotehrCustomSlice{}
	cs = as
	as = cs

	// cs = ais not allowed
	cs = customSlice(is)
	acs = anotehrCustomSlice(cs)
	fmt.Println(reflect.TypeOf(acs)) //[]
}
