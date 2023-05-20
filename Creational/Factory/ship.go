package main

type Ship struct {
	Transport
}

func newShip() TransportInter {
	return &Ship{
		Transport: Transport{
			name:  "Ship",
			power: 10,
		},
	}
}
