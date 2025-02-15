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
	me := point{ }

	you := point{
		x: 0,
		y: 0,
	}

	res := me == you
	fmt.Printf("res: %T", res)

	log.SetPrefix("> ")
	log.SetFlags(0)
	msg, err := greeting.Okay("hemlo")
	if is_err(err) {
		fmt.Println("bad", err)
	}

	fmt.Printf("%s", msg)
}

func is_err(err error) bool {
	if err != nil {
		log.Print("language!")
		return true
	}
	return false
}
