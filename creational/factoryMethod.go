package creational

import (
	"errors"
	"fmt"
)

// transport example
// this pattern is used to create objects, hiding the specific creation logic
// and providing factory methods to the client

type ITransport interface {
	setName(name string)
	setCapacity(capacity int)
	getName() string
	getCapacity() int
}

type Transport struct {
	name     string
	capacity int
}

func (t *Transport) setName(name string) {
	t.name = name
}

func (t *Transport) setCapacity(capacity int) {
	t.capacity = capacity
}

func (t *Transport) getName() string {
	return t.name
}

func (t *Transport) getCapacity() int {
	return t.capacity
}

type Motorbike struct {
	Transport
}

func newMotorbike() ITransport {
	return &Motorbike{
		Transport{
			name:     "Suzuki Hayabusa",
			capacity: 2,
		},
	}
}

type Car struct {
	Transport
}

func newCar() ITransport {
	return &Car{
		Transport{
			name:     "Hyundai Sonata",
			capacity: 5,
		},
	}
}

type Bus struct {
	Transport
}

func newBus() ITransport {
	return &Bus{
		Transport{
			name:     "Mercedes",
			capacity: 37,
		},
	}
}

func getTransport(transportType string) (ITransport, error) {
	if transportType == "motorbike" {
		return newMotorbike(), nil
	}

	if transportType == "car" {
		return newCar(), nil
	}

	if transportType == "bus" {
		return newBus(), nil
	}

	return nil, errors.New("wrong transport type passed")
}

func RunFactoryMethod() {
	motorbike, _ := getTransport("motorbike")
	fmt.Printf("transport name is: %s\n", motorbike.getName())
	fmt.Printf("transport capacity is: %d\n", motorbike.getCapacity())

	car, _ := getTransport("car")
	fmt.Printf("transport name is: %s\n", car.getName())
	fmt.Printf("transport capacity is: %d\n", car.getCapacity())

	bus, _ := getTransport("bus")
	fmt.Printf("transport name is: %s\n", bus.getName())
	fmt.Printf("transport capacity is: %d\n", bus.getCapacity())
}
