package simplefactory

import (
	"testing"
)

func TestNewFood(t *testing.T) {
	Restaurant("rice").Cook()
	Restaurant("noddle").Cook()
}
