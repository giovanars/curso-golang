package arquitetura

import "testing"
import "runtime"

func TestDependete(t *testing.T) {
	if runtime.GOARCH == "amd64"{
		t.Skip("Não funciona na arquitetura amd64")
	}
}