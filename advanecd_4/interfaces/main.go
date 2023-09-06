package main

import "fmt"

// type alias
type Helper = interface {
	Help() string
}

type HelpString string

func (h HelpString) Help() string {
	return string(h)
}

type UnHelpString struct {
}

func (uhs *UnHelpString) Help() string {
	return "I can not help you"
}

// compile time check
var _ = Helper(HelpString(""))

func main() {
	fmt.Println(HelpString("Hey").Help())
	//fmt.Println(UnHelpString{}.Help())

	for _, v := range []Helper{HelpString(""), &UnHelpString{}} {
		fmt.Println(v.Help())
	}
}
