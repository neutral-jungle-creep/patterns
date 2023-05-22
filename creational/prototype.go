package creational

import "fmt"

// box example
// this example uses custom interface Clonable with function clone & helper function print
// functions for struct Box works recursively which allows to bypass all nesting levels

type Cloneable interface {
	clone() Cloneable
	print(indent string)
}

type Toy struct {
	name string
}

func (t *Toy) clone() Cloneable {
	return &Toy{
		name: t.name + "_clone",
	}
}

func (t *Toy) print(indent string) {
	fmt.Println(indent + t.name)
}

type Box struct {
	content []Cloneable
	name    string
}

func (b *Box) clone() Cloneable {
	var cloneBox = &Box{name: b.name + "_clone"}

	var cloneContent []Cloneable
	for _, i := range b.content {
		copy := i.clone()
		cloneContent = append(cloneContent, copy)
	}

	cloneBox.content = cloneContent
	return cloneBox
}

func (b *Box) print(indent string) {
	fmt.Println(indent + b.name)

	for _, c := range b.content {
		c.print(indent + indent)
	}
}

func RunPrototype() {
	var ball = &Toy{name: "ball"}
	var dall = &Toy{name: "dall"}

	var littleBox = &Box{
		name:    "little box",
		content: []Cloneable{ball, dall},
	}
	var bigBox = &Box{
		name:    "big box",
		content: []Cloneable{littleBox},
	}

	cloneBox := bigBox.clone()
	bigBox.print("--")
	cloneBox.print("--")
}
