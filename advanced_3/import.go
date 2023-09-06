package main

import (
	"log"
	_ "log"  // will only call init if any
	l "log"  // only useful if we've package which has the same interface (exported identifiers) as other imported package
	. "math" // will break if 2 packages have common exported ientifiers

	"github.com/golang/glog"
)

func main() {
	log.Println("Log Entry: 1")
	l.Println("Log Entry: 2")
	l.Printf("Log Entry: %d", int(Floor(3.3)))
	glog.Warning("warning before exit")
}
