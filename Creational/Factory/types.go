package main

type TransportInter interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Transport struct {
	name  string
	power int
}

func (g *Transport) setName(name string) {
	g.name = name
}

func (g *Transport) getName() string {
	return g.name
}

func (g *Transport) setPower(power int) {
	g.power = power
}

func (g *Transport) getPower() int {
	return g.power
}
