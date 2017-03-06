package agent

import "github.com/maksadbek/dpipe"

type Gatherer struct {
	hotelsc chan dpipe.Hotel
}

func (g *Gatherer) Write(h dpipe.Hotel) {
	g.hotelsc <- h
}
