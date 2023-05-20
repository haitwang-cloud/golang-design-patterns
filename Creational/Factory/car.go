package main

type Car struct {
    Transport
}

func newCar() TransportInter {
    return &Car{
        Transport: Transport{
            name:  "Car",
            power: 4,
        },
    }
}