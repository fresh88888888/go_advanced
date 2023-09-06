package main

import (
	"fmt"

	"umbrella.github.com/advanced_go/advanced_5/dev-data-structure/set"
)

func main() {
	var s = set.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)

	fmt.Println(s.Size())
	s.Add(1)
	s.Remove(6)
	fmt.Println(s.Size())
	fmt.Println(s.IsElementOf(3))
	fmt.Println(s.Values())
}
