package main

import (
	"fmt"
)

type Autobot struct {
	AlphaTrion   string
	OptimusPrime string
	UltraMagnus  string
	Bumblebee    string
	Sideswipe    string
	Wheeljack    string
	Skyfire      string
	Moonracer    string
	Elita        string
}

type Centrist struct {
}

type Decepticon struct {
	Megatron       string
	Starscream     string
	Thundercracker string
	Soundwave      string
	Menasor        string
	Laserbeak      string
	Swindle        string
	Shockwave      string
}

type Transformer interface {
	transform() string
}

func (a *Autobot) transform() {
	a.OptimusPrime = fmt.Sprintf("Autobot boss: Optimus Prime.")
}

func (c *Centrist) transform() {
}

func (d *Decepticon) transform() {
	d.Megatron = fmt.Sprintf("Decepticon boss: Megatron.")
}

func main() {
	fmt.Println("Welcome to Cybertron!")
	a := &Autobot{}
	a.transform()
	fmt.Println(a)

	d := &Decepticon{}
	d.transform()
	fmt.Println(d)
}
