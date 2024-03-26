package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		expectedArea := float64(120)
		receivedArea := ret.Area()

		if receivedArea != expectedArea {
			t.Fatalf("A area %f é diferente da esperada %f", receivedArea, expectedArea)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if receivedArea != expectedArea {
			t.Fatalf("A area %f é diferente da esperada %f", receivedArea, expectedArea)
		}
	})
}
