package creational

import (
	"errors"
	"fmt"
)

// appliances example
// this pattern creates factories of different brands
// when using the same factory, all items are guaranteed to be the same brand

type IFactory interface {
	makeMobil() IMobil
	makeCharger() ICharger
}

func GetFactory(brand string) (IFactory, error) {
	if brand == "apple" {
		return &Apple{}, nil
	}

	if brand == "nokia" {
		return &Nokia{}, nil
	}

	return nil, errors.New("wrong mobile type passed")
}

type Apple struct {
}

func (a *Apple) makeMobil() IMobil {
	return &AppleMobil{
		Mobil{
			cost: 60000,
		},
	}
}

func (a *Apple) makeCharger() ICharger {
	return &AppleCharger{
		Charger{
			cost: 3000,
		},
	}
}

type Nokia struct {
}

func (n *Nokia) makeMobil() IMobil {
	return &NokiaMobil{
		Mobil{
			cost: 15000,
		},
	}
}

func (n *Nokia) makeCharger() ICharger {
	return &NokiaCharger{
		Charger{
			cost: 500,
		},
	}
}

type IMobil interface {
	setCost(cost int)
	getCost() int
}

type Mobil struct {
	cost int
}

func (m *Mobil) setCost(cost int) {
	m.cost = cost
}

func (m *Mobil) getCost() int {
	return m.cost
}

type AppleMobil struct {
	Mobil
}

type NokiaMobil struct {
	Mobil
}

type ICharger interface {
	setCost(cost int)
	getCost() int
}

type Charger struct {
	cost int
}

func (c *Charger) setCost(cost int) {
	c.cost = cost
}

func (c *Charger) getCost() int {
	return c.cost
}

type AppleCharger struct {
	Charger
}

type NokiaCharger struct {
	Charger
}

func RunAbstractFactory() {
	appleFactory, _ := GetFactory("apple")

	iphone := appleFactory.makeMobil()
	lightning := appleFactory.makeCharger()

	fmt.Printf("apple mobile costs: %d\n", iphone.getCost())
	fmt.Printf("apple charger costs: %d\n", lightning.getCost())

	nokiaFactory, _ := GetFactory("nokia")

	nAndroid := nokiaFactory.makeMobil()
	microUSB := nokiaFactory.makeCharger()

	fmt.Printf("nokia mobile costs: %d\n", nAndroid.getCost())
	fmt.Printf("nokia charger costs: %d\n", microUSB.getCost())

}
