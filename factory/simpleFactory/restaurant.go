package simplefactory

import "fmt"

type Food interface {
	Cook()
}

type Rice struct{}

func (r *Rice) Cook() {
	fmt.Println("Rice is cooking.")
}

type Noddle struct{}

func (n *Noddle) Cook() {
	fmt.Println("Noddle is cooking.")
}

func Restaurant(name string) Food {
	switch name {
	case "rice":
		return &Rice{}
	case "noddle":
		return &Noddle{}
	}
	return nil
}
