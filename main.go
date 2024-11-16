package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"git.sriss.uz/shared/shared_service/logger"
)

type (
	Bool bool
)

type String string

func (s String) Println() {
	log.Println(s)
}

type Ints []int

type Int int

func main() {

	f, _ := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	l := logger.New()

	l.Info("Hello World!")
	l.Debug("Hello World!")
	l.Error("Hello World!")
	l.Trace("Hello World!")
	l.Println("Hello World!")
}

// type Point struct {
// 	X float64
// 	Y float64
// 	S string
// }

// p := Point{}

// n, err := fmt.Sscanf(`(1,2.2)`, "(%f,%f)", &p.X, &p.Y)
// log.Println(n)
// log.Println(err)
// log.Println(p)

func exampleFunction() {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		function := runtime.FuncForPC(pc)
		fmt.Println(function.FileLine(pc))
		fmt.Printf("Function: %s\nFile: %s\nLine: %d\n", function.Name(), file, line)
	}
}
