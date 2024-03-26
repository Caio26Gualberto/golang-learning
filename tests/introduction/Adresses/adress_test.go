//Teste de Unidade

package adresses

import (
	"testing"
)

type TestScenario struct {
	adressInserted string
	expectedReturn string
}

func TestAdressType(t *testing.T) {
	tests := []TestScenario{
		{"Street ABC", "Street"},
		{"Avenue Paulista", "Avenue"},
		{"Road 66", "Road"},
		{"Highway 666", "Highway"},
		{"Estrada", "Invalid type!"},
	}

	for _, scenario := range tests {
		adressTypeReceived := AdressType(scenario.adressInserted)

		if adressTypeReceived != scenario.expectedReturn {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", adressTypeReceived, scenario.expectedReturn)
		}
	}
}
