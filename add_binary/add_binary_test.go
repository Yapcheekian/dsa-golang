package addbinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	result := AddBinary(a, b)
	assert.Equal(t, "100", result)

	a = "1010"
	b = "1011"
	result = AddBinary(a, b)
	assert.Equal(t, "10101", result)

	a = "0"
	b = "0"
	result = AddBinary(a, b)
	assert.Equal(t, "0", result)
}
