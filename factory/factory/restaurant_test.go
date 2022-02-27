package factory

import (
	"testing"
)

func TestRestaurant(t *testing.T) {
	Restaurant("rice").NewFood().Cook()
	Restaurant("noddle").NewFood().Cook()
}
