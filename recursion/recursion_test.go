package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPrintInteger(t *testing.T) {
// 	PrintInteger(-1496)
// 	fmt.Println("")
// }

// func TestMoveTower(t *testing.T) {
// 	MoveTower(4, "A", "B", "C")
// }

// func TestListPermutations(t *testing.T) {
// 	ListPermutations("ABCD")
// fmt.Println("")
// }

func TestReverseString(t *testing.T) {
	s := ReverseString("abcdefg")
	assert.Equal(t, "gfedcba", s)
}

func TestCountX(t *testing.T) {
	s := CountX("xxgfrtxthxooxxr")
	assert.Equal(t, 6, s)
}
