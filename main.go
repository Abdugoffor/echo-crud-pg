package main

import (
	"log"
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
