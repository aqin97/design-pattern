package abstractfactory

import "fmt"

type HotPot interface {
	Cook()
}

type SpicyHotPot struct{}

func (s *SpicyHotPot) Cook() {
	fmt.Println("Spicy hotpot has been cooked")
}

type NormalHotPot struct{}

func (n *NormalHotPot) Cook() {
	fmt.Println("Normal hotpot has been cooked")
}

type FriedFood interface {
	Fry()
}

type FranchFries struct{}

func (f *FranchFries) Fry() {
	fmt.Println("Franch Fries has been fried")
}

type Chips struct{}

func (c *Chips) Fry() {
	fmt.Println("Chips has been fried")
}

type Restaurant interface {
	CreateHotPot(string) HotPot
	CreateFriedFood(string) FriedFood
}

type RestaurantA struct{}

func (ra *RestaurantA) CreateHotPot(name string) HotPot {
	switch name {
	case "spicy":
		return &SpicyHotPot{}
	case "normal":
		return &NormalHotPot{}
	}

	return nil
}

func (ra *RestaurantA) CreateFriedFood(name string) FriedFood {
	switch name {
	case "fries":
		return &FranchFries{}
	case "chips":
		return &Chips{}
	}

	return nil
}

type RestaurantB struct{}

func (ra *RestaurantB) CreateHotPot(name string) HotPot {
	switch name {
	case "spicy":
		return &SpicyHotPot{}
	case "normal":
		return &NormalHotPot{}
	}

	return nil
}

func (ra *RestaurantB) CreateFriedFood(name string) FriedFood {
	switch name {
	case "fries":
		return &FranchFries{}
	case "chips":
		return &Chips{}
	}

	return nil
}

func GetRestaurant(name string) Restaurant {
	switch name {
	case "a":
		return &RestaurantA{}
	case "b":
		return &RestaurantB{}
	}
	return nil
}
