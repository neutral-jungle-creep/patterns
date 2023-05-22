package creational

import "fmt"

// house example
// this example uses universal function getBuilder for build different houses with difficult structs

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "wooden" {
		return newWoodenBuilder()
	}

	if builderType == "stone" {
		return newStoneBuilder()
	}
	return nil
}

type WoodenBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newWoodenBuilder() *WoodenBuilder {
	return &WoodenBuilder{}
}

func (b *WoodenBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *WoodenBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *WoodenBuilder) setNumFloor() {
	b.floor = 2
}

func (b *WoodenBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type StoneBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newStoneBuilder() *StoneBuilder {
	return &StoneBuilder{}
}

func (b *StoneBuilder) setWindowType() {
	b.windowType = "Stone Window"
}

func (b *StoneBuilder) setDoorType() {
	b.doorType = "Stone Door"
}

func (b *StoneBuilder) setNumFloor() {
	b.floor = 5
}

func (b *StoneBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

func RunBuilder() {
	builder := getBuilder("wooden")

	builder.setWindowType()
	builder.setDoorType()
	builder.setNumFloor()
	woodenHouse := builder.getHouse()

	builder = getBuilder("stone")

	builder.setWindowType()
	builder.setDoorType()
	builder.setNumFloor()
	stoneHouse := builder.getHouse()

	fmt.Println("Wooden House - Window Type:", woodenHouse.windowType)
	fmt.Println("Wooden House - Door Type:", woodenHouse.doorType)
	fmt.Println("Wooden House - Number of Floors:", woodenHouse.floor)

	fmt.Println("Stone House - Window Type:", stoneHouse.windowType)
	fmt.Println("Stone House - Door Type:", stoneHouse.doorType)
	fmt.Println("Stone House - Number of Floors:", stoneHouse.floor)
}
