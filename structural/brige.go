package structural

import "fmt"

// share example
// this example uses various combinations of colored shapes
// interface Painter built into the Share's structs and its method is called in the Draw method

type Share interface {
	Draw()
	SetColor(Painter)
}

type Circle struct {
	painter Painter
}

func (c *Circle) Draw() {
	c.painter.Paint("circle")
}

func (c *Circle) SetColor(p Painter) {
	c.painter = p
}

type Square struct {
	painter Painter
}

func (s *Square) Draw() {
	s.painter.Paint("square")
}

func (s *Square) SetColor(p Painter) {
	s.painter = p
}

type Painter interface {
	Paint(shareName string)
}

type Red struct {
}

func (r *Red) Paint(shareName string) {
	fmt.Printf("%s is red\n", shareName)
}

type Green struct {
}

func (g *Green) Paint(shareName string) {
	fmt.Printf("%s is green\n", shareName)
}

func RunBridge() {
	circle := &Circle{}
	squire := &Square{}

	redColor := &Red{}
	greenColor := &Green{}

	circle.SetColor(redColor)
	circle.Draw()
	circle.SetColor(greenColor)
	circle.Draw()

	squire.SetColor(greenColor)
	squire.Draw()
}
