package main

import "fmt"

const (
	carType  = "car"
	shipType = "ship"
)

func genTransport(TransportType string) (TransportInter, error) {
	switch TransportType {
	case carType:
		return newCar(), nil
	case shipType:
		return newShip(), nil
	default:
		return nil, fmt.Errorf("transport of type %s not recognized", TransportType)
	}
}

func genDetails(g TransportInter) {
	fmt.Printf("Name: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
func setDetails(g TransportInter) {
	g.setPower(g.getPower() * 2)
}

func main() {
	car, _ := genTransport(carType)
	ship, _ := genTransport(shipType)
	genDetails(car)
	genDetails(ship)
	setDetails(car)
	setDetails(ship)
	genDetails(car)
	genDetails(ship)
}
