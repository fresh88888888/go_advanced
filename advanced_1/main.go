package main

import (
	"fmt"
	"log"
	"math"
	"sync"
)

var atom int

const monad = "monad"

// known when init is called
func init() {
	atom = 1
}

// function declaration
func calculated(input float64) (result float64, err error) {
	return math.Ceil(input), nil
}

// type definations and aliases
type ByteSlice []byte
type Bytes = ByteSlice

func between(from, to int) []int {
	var result []int
	if from > to {
		return []int{}
	} else {
		for i := 0; i < to; i++ {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// some basic print flags
	fmt.Printf("%d\n", atom)
	fmt.Printf("%s\n", monad)

	// short declaration and assignment
	string_slice := [3]string{"1", "2", "3"}
	// basic slice operations
	fmt.Println(len(string_slice))

	// checking for errors
	_, err := calculated(1.23)
	if err != nil {
		log.Fatal(err)
	}

	// infinite loop
	counter := 0
	for {
		if counter > 1000 {
			break
		} else {
			counter++
		}
	}

	// type switches
	var boolRef interface{}
	boolean := true
	boolRef = &boolean

	switch t := boolRef.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolen %t\n", *t)
	case *int:
		fmt.Printf("pointer to int %d\n", *t)
	}

	// defer calls as lifo order
	defer fmt.Printf("%d", 1)
	defer fmt.Printf("%d", 2)

	lock := new(sync.Mutex)
	lock.Lock()
	defer lock.Unlock()

	// interfaces and export rules
	type Iterator interface {
		Next() interface{}
		HasNext() bool
	}

	// declare struct
	type Guard struct {
	}

	// declare channels
	ch := make(chan Guard, 1)

	go func(chan Guard) {
		fmt.Printf("%T\n", ch)
		ch <- Guard{}
		close(ch)
	}(ch)

	<-ch

	for i := range between(0, 10) {
		switch i % 5 {
		case 1:
			fmt.Println("fizz")
		case 2:
			fmt.Println("bazz")
		case 3:
			fmt.Println("gizz")
		default:
			fmt.Println("fizzbazz")
		}
	}
}
