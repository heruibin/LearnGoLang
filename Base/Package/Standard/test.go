package standard

import (
	"testing"
)

func TestCToF(t *testing.T) {
	var c Celsius = 100
	t.Log(CToF(c))
}
