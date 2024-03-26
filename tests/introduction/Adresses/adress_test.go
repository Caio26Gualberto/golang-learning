//Teste de Unidade

package adresses

import (
	"testing"
)

func TestAdressType(t *testing.T) {
	addresForTest := "Avenue Paulista"

	addresTypeExpected := "Avenue"

	addresTypeReceived := AdressType(addresForTest)

	if addresTypeReceived != addresTypeExpected {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", addresTypeExpected, addresTypeReceived)
	}
}
