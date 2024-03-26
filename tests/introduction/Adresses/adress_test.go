//Teste de Unidade

package adresses // go test -> go test ./... -> go test -v -> go test --cover -> go test --coverprofile (OutputFile: any text of your choice here) ->
//go tool cover --func=(Here you put your output file name) -> go tool cover --html=(Here you put your output file name)

import (
	"testing"
)

type TestScenario struct {
	adressInserted string
	expectedReturn string
}

func TestAdressType(t *testing.T) {
	t.Parallel() // Run Parallel test
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

func TestAny(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Test Fail !")
	}
}
