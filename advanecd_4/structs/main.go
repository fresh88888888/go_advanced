package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age  int
}

func main() {
	alex := User{}
	fmt.Println(alex.age)
	alexP := &alex
	fmt.Println(alexP)

	var worker = struct {
		User
		salary int
	}{
		User:   alex,
		salary: 100000,
	}

	var anotherWorker = struct {
		User
		salary int
	}{
		User: struct {
			name string
			age  int
		}{
			name: "",
			age:  0,
		},
		salary: 100000,
	}

	fmt.Println(worker.salary)
	fmt.Println(worker.name)
	fmt.Println(worker.age)

	fmt.Println(worker == anotherWorker) //true

	a := struct {
		name string
		age  int
	}{"", 0}
	fmt.Println(reflect.DeepEqual(a, alex)) // false
}
