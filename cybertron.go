package cybertron

import (
	"fmt"
)

const (
	A = "Autobots"
	D = "Decepticons"
)

const (
	SEASON1 = iota + 1
	SEASON2
	SEASON3
)

type Transformers struct {
	Value    float64
	Deformed string
}

type OptimusPrime struct {
	Transformers
}

type UltraMagnus struct {
	Transformers
}

type Bumblebee struct {
	Transformers
}

type Autobot struct {
	OptimusPrime
	UltraMagnus
	Bumblebee
	Season int
	//Sideswipe    string
	//Wheeljack    string
	//Skyfire      string
	//Moonracer    string
	//Elita        string
}

type DoubleDealer struct {
	Transformers
	factions string
}

type Centrist struct {
	DoubleDealer
	Season int
}

type Megatron struct {
	Transformers
}

type Starscream struct {
	Transformers
}

type Thundercracker struct {
	Transformers
}

type Decepticon struct {
	Megatron
	Starscream
	Thundercracker
	Season int
	//Soundwave      string
	//Menasor        string
	//Laserbeak      string
	//Swindle        string
	//Shockwave      string
}

type Transformer interface {
	transform(season ...int)
	init(season int)
}

// TODO: force allocation algorithm for all autobots
func (a *Autobot) init(season ...int) {
	if len(season) > 0 {
		a.Season = season[0]
	}
	switch a.Season {
	case SEASON1:
		a.OptimusPrime.Value = 80.0
	case SEASON2:
		a.OptimusPrime.Value = 85.0
	case SEASON3:
		a.OptimusPrime.Value = 90.0
	default:
	}
}

func (a *Autobot) transform(season ...int) {
	a.OptimusPrime.Deformed = fmt.Sprint("Peterbilt")
}

func (c *Centrist) init(season ...int) {
	if len(season) > 0 {
		c.Season = season[0]
	}
	switch c.Season {
	case SEASON1:
	case SEASON2:
	case SEASON3:
	default:
	}

}

func (c *Centrist) transform(season ...int) {
	// TODO: ensure the deformed of duble-dealer
	if c.DoubleDealer.factions == A {
		c.DoubleDealer.Deformed = ""
	}
	c.DoubleDealer.Deformed = ""
}

func (d *Decepticon) init(season ...int) {
	if len(season) > 0 {
		d.Season = season[0]
	}
	switch d.Season {
	default:
	}

}

func (d *Decepticon) transform(season ...int) {
	if len(season) > 0 {
		d.Season = season[0]
	}
	switch d.Season {
	case SEASON1:
		d.Megatron.Deformed = fmt.Sprintf("Warcraft")
	case SEASON2:
		d.Megatron.Deformed = fmt.Sprintf("Panzer")
	case SEASON3:
		d.Megatron.Deformed = fmt.Sprintf("OilTank")
	default:
	}
}
