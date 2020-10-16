package cybertron

import (
	"fmt"
)

const (
	SEASON1 = iota + 1
	SEASON2
	SEASON3
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
	season         int
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
	transform(season ...int)
}

func (a *Autobot) transform(season ...int) {
	a.OptimusPrime = fmt.Sprint("Peterbilt")
}

func (c *Centrist) transform(season ...int) {
}

func (d *Decepticon) transform(season ...int) {
	if len(season) > 0 {
		d.season = season[0]
	}
	switch d.season {
	case SEASON1:
		d.Megatron = fmt.Sprintf("Warcraft")
	case SEASON2:
		d.Megatron = fmt.Sprintf("Panzer")
	case SEASON3:
		d.Megatron = fmt.Sprintf("OilTank")
	}
}
