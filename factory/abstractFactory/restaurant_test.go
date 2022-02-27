package abstractfactory

import (
	"testing"
)

func TestGetFood(t *testing.T) {
	GetRestaurant("a").CreateHotPot("spicy").Cook()
	GetRestaurant("a").CreateHotPot("normal").Cook()
	GetRestaurant("a").CreateFriedFood("fries").Fry()
	GetRestaurant("a").CreateFriedFood("chips").Fry()

	GetRestaurant("b").CreateHotPot("spicy").Cook()
	GetRestaurant("b").CreateHotPot("normal").Cook()
	GetRestaurant("b").CreateFriedFood("fries").Fry()
	GetRestaurant("b").CreateFriedFood("chips").Fry()
}
