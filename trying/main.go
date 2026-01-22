package main

import (
	"fmt"
	"greeting/fine"
	"log"
)

type point struct {
	x int
	y int
	z int 
}

func (p *point) SetX(x int) {
	p.x = x
}

func (p *point) SetY(y int) {
	p.y = y
}

func (p *point) SetZ(z int) {
	p.z = z
}

func main() {
	log.SetPrefix("> ")
	log.SetFlags(0 | log.Lshortfile)
	me := point{ }

	you := point{
		x: 0,
		y: 1,
	}

	res := me == you
	fmt.Printf("res: (%T) %v\n", res, res)
	fmt.Printf("arg: (%T) %v\n", you, you)

	msg, err := greeting.Okay("")
	if is_err(err) {
		fmt.Println("bad", err)
	}

	fmt.Printf("%s\n", msg)
}

func is_err(err error) bool {
	if err != nil {
		log.Print("language!")
		return true
	}
	return false
}
