package factory

import "fmt"

type Food interface {
	Cook()
}

type Kitchen interface {
	NewFood() Food
}

type Rice struct{}

func (r *Rice) Cook() {
	fmt.Println("Rice is cooking.")
}

type Noddle struct{}

func (n *Noddle) Cook() {
	fmt.Println("Noddle is cooking.")
}

type RiceKitchen struct{}

func (r *RiceKitchen) NewFood() Food {
	return &Rice{}
}

type NoddleKitchen struct{}

func (n *NoddleKitchen) NewFood() Food {
	return &Noddle{}
}

func Restaurant(name string) Kitchen {
	switch name {
	case "rice":
		return &RiceKitchen{}
	case "noddle":
		return &NoddleKitchen{}
	}

	return nil
}
