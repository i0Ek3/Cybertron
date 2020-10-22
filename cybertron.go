package cybertron

import (
	"fmt"
)

// defines transformers type
const (
	A = "Autobots"
	D = "Decepticons"
)

// defines transformers' season
const (
	SEASON1 = iota + 1
	SEASON2
	SEASON3
)

// Transformers defines the transformers'
// power and what they can transformered
type Transformers struct {
	Value    float64
	Deformed string
}

// qtz
type OptimusPrime struct {
	Transformers
}

// ttx
type UltraMagnus struct {
	Transformers
}

// dhf
type Bumblebee struct {
	Transformers
}

// hp
type Sideswipe struct {
	Transformers
}

// qjd
type Wheeljack struct {
	Transformers
}

// th
type Skyfire struct {
	Transformers
}

// alt
type Elita struct {
	Transformers
}

// Autobot defines all autobots
type Autobot struct {
	OptimusPrime
	UltraMagnus
	Bumblebee
	Sideswipe
	Wheeljack
	Skyfire
	Elita
	Season int
	isDD   bool
}

// smr
type DoubleDealer struct {
	Transformers
	factions string
	trueName string
}

// Centrist defines double-dealer
type Centrist struct {
	DoubleDealer
	Season int
	isDD   bool
}

// wzt
type Megatron struct {
	Transformers
}

// hzz
type Starscream struct {
	Transformers
}

// ld
type Thundercracker struct {
	Transformers
}

// zdb
type Shockwave struct {
	Transformers
}

// sb
type Soundwave struct {
	Transformers
}

// zp
type Swindle struct {
	Transformers
}

// jgn
type Menasor struct {
	Transformers
}

// Decepticon defines all decepticons
type Decepticon struct {
	Megatron
	Starscream
	Thundercracker
	Shockwave
	Soundwave
	Swindle
	Menasor
	Season int
	isDD   bool
}

// Transformer defines transformers' interface
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

	a.isDD = false
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

	c.isDD = true
}

func (c *Centrist) transform(season ...int) {
	if c.isDD {
		c.DoubleDealer.Deformed = "Pontiac"
		if c.DoubleDealer.factions == A {
			c.DoubleDealer.trueName = "CounterPunch"
		}
		c.DoubleDealer.trueName = "Punch"
	}
}

func (d *Decepticon) init(season ...int) {
	if len(season) > 0 {
		d.Season = season[0]
	}
	switch d.Season {
	default:
	}

	d.isDD = false
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
