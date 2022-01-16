package satisfiabilityofequailityequation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquationsPossible(t *testing.T) {
	equations := []string{"a==b", "b!=a"}
	result := EquationsPossible(equations)
	assert.Equal(t, false, result)

	equations = []string{"a==b", "b==c", "c==d", "d!=a"}
	result = EquationsPossible(equations)
	assert.Equal(t, false, result)

	equations = []string{"a==b", "b==a"}
	result = EquationsPossible(equations)
	assert.Equal(t, true, result)
}
